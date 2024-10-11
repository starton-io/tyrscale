package validation

import (
	"reflect"
	"regexp"
	"strings"

	enLocales "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

// Option validation option
type Option interface {
	apply(*option)
}

// option implement
type option struct {
	validator *validator.Validate
	uni       *ut.UniversalTranslator
	trans     *ut.Translator
}

type optionFn func(*option)

func (optFn optionFn) apply(opt *option) {
	optFn(opt)
}

// WithValidator set validator
func WithValidator(v *validator.Validate) Option {
	return optionFn(func(opt *option) {
		opt.validator = v
	})
}

// WithUniversalTranslator set UniversalTranslator
func WithUniversalTranslator(uni *ut.UniversalTranslator) Option {
	return optionFn(func(opt *option) {
		opt.uni = uni
	})
}

// WithTranslator set Translator
func WithTranslator(trans *ut.Translator) Option {
	return optionFn(func(opt *option) {
		opt.trans = trans
	})
}

func getDefaultOption() *option {
	v := validator.New()
	// custom validator for regexp
	err := v.RegisterValidation("regexp", Regexp)
	if err != nil {
		panic(err)
	}
	translator := enLocales.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		return nil
	}

	err = v.RegisterTranslation("regexp", trans, func(ut ut.Translator) error {
		return ut.Add("regexp", "{0} is not respect this regex: {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("regexp", fe.Field(), fe.Param())
		return t
	})
	if err != nil {
		return nil
	}

	if err := enTranslations.RegisterDefaultTranslations(v, trans); err != nil {
		return nil
	}

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tag := fld.Tag.Get("json")
		if tag == "" {
			tag = fld.Tag.Get("query")
		}
		if tag == "" {
			return fld.Name
		}

		name := strings.SplitN(tag, ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	return &option{
		validator: v,
		uni:       uni,
		trans:     &trans,
	}
}

func getOption(opts ...Option) *option {
	opt := getDefaultOption()
	for _, o := range opts {
		o.apply(opt)
	}

	return opt
}

func Regexp(fl validator.FieldLevel) bool {
	// Replace all occurrences of \\ with \
	param := strings.ReplaceAll(fl.Param(), "\\\\", "\\")
	re, err := regexp.Compile(param)
	if err != nil {
		return false
	}
	return re.MatchString(fl.Field().String())
}
