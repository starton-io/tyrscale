package validation

// Validation is an interface for validating structs
type Validation interface {
	ValidateStruct(s interface{}) error
}
