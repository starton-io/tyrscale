package logger

import (
	"context"
	"fmt"
	"log/slog"
	"testing"

	slogzap "github.com/samber/slog-zap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

var (
	logger  Logger
	sLogger *slog.Logger
)

type PrintfLogger interface {
	Printf(string, ...any)
}

type Logger interface {
	PrintfLogger
	Debug(msg string, fields ...interface{})
	Debugf(msg string, args ...interface{})
	Info(msg string, fields ...interface{})
	Infof(msg string, args ...interface{})
	Warn(msg string, fields ...interface{})
	Warnf(msg string, args ...interface{})
	Error(msg string, fields ...interface{})
	Errorf(msg string, args ...interface{})
	Fatal(msg string, fields ...interface{})
	Fatalf(msg string, args ...interface{})
	Logf(msg string, args ...interface{})
}

type SLogger interface {
	DebugContext(ctx context.Context, msg string, fields ...interface{})
	InfoContext(ctx context.Context, msg string, fields ...interface{})
	WarnContext(ctx context.Context, msg string, fields ...interface{})
	ErrorContext(ctx context.Context, msg string, fields ...interface{})
}

type ZapLogger struct {
	Logger       *zap.Logger
	loggerConfig zap.Config
}

type optionFunc func(*ZapLogger)

func InitLogger(opts ...optionFunc) error {
	if logger != nil {
		return nil
	}
	zapLogger, err := NewZapLogger(opts...)
	if err != nil {
		return err
	}
	logger = zapLogger
	return nil
}

func NewZapLogger(opts ...optionFunc) (*ZapLogger, error) {
	if logger != nil {
		return logger.(*ZapLogger), nil
	}
	loggerConfig := zap.NewProductionConfig()
	loggerZap := &ZapLogger{loggerConfig: loggerConfig}
	for _, opt := range opts {
		opt(loggerZap)
	}
	var err error
	loggerZap.Logger, err = loggerZap.loggerConfig.Build()
	if err != nil {
		return nil, err
	}
	logger = loggerZap
	return loggerZap, nil
}

type LevelAdapter struct {
	ZapLevel zapcore.Level
}

func (l LevelAdapter) Level() slog.Level {
	switch l.ZapLevel {
	case zapcore.DebugLevel:
		return slog.LevelDebug
	case zapcore.InfoLevel:
		return slog.LevelInfo
	case zapcore.WarnLevel:
		return slog.LevelWarn
	case zapcore.ErrorLevel:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

type OptSLogger struct {
	AttrFromCtx []func(ctx context.Context) []slog.Attr
	ZapLevel    zapcore.Level
}

func NewSLoggerFromZap(zapLogger *zap.Logger, opts *OptSLogger) *slog.Logger {
	if sLogger != nil {
		return sLogger
	}
	slogger := slog.New(slogzap.Option{Logger: zapLogger, Level: LevelAdapter{ZapLevel: opts.ZapLevel}, AttrFromContext: opts.AttrFromCtx}.NewZapHandler())
	sLogger = slogger
	return slogger
}

func DebugContext(ctx context.Context, msg string, fields ...interface{}) {
	sLogger.DebugContext(ctx, msg, fields...)
}

func InfoContext(ctx context.Context, msg string, fields ...interface{}) {
	sLogger.InfoContext(ctx, msg, fields...)
}

func WarnContext(ctx context.Context, msg string, fields ...interface{}) {
	sLogger.WarnContext(ctx, msg, fields...)
}

func ErrorContext(ctx context.Context, msg string, fields ...interface{}) {
	sLogger.ErrorContext(ctx, msg, fields...)
}

func NewZapLoggerForTest(t *testing.T) Logger {
	return &ZapLogger{
		Logger: zaptest.NewLogger(t),
	}
}

func WithLevel(level zapcore.Level) optionFunc {
	return func(zl *ZapLogger) {
		zl.loggerConfig.Level = zap.NewAtomicLevelAt(level)
	}
}
func WithEncodeTime(timeKey string, timeEncoder zapcore.TimeEncoder) optionFunc {
	return func(zl *ZapLogger) {
		zl.loggerConfig.EncoderConfig.TimeKey = timeKey
		zl.loggerConfig.EncoderConfig.EncodeTime = timeEncoder
	}
}

func GetLogger() Logger {
	return logger
}

func GetSLogger() SLogger {
	return sLogger
}

func Debug(msg string, fields ...interface{}) {
	logger.Debug(msg, fields...)
}

func Debugf(msg string, fields ...interface{}) {
	logger.Debugf(msg, fields...)
}

func Info(msg string, fields ...interface{}) {
	logger.Info(msg, fields...)
}

func Infof(msg string, fields ...interface{}) {
	logger.Infof(msg, fields...)
}

func Warn(msg string, fields ...interface{}) {
	logger.Warn(msg, fields...)
}

func Warnf(msg string, fields ...interface{}) {
	logger.Warnf(msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	logger.Error(msg, fields...)
}
func Errorf(msg string, fields ...interface{}) {
	logger.Errorf(msg, fields...)
}

func Fatal(msg string, fields ...interface{}) {
	logger.Fatal(msg, fields...)
}

func Fatalf(msg string, fields ...interface{}) {
	logger.Fatalf(msg, fields...)
}

func (l *ZapLogger) Debug(msg string, fields ...interface{}) {
	l.Logger.Sugar().Debugw(msg, fields...)
}

func (l *ZapLogger) Debugf(msg string, args ...interface{}) {
	l.Logger.Sugar().Debugf(msg, args...)
}

func (l *ZapLogger) Info(msg string, fields ...interface{}) {
	l.Logger.Sugar().Infow(msg, fields...)
}

func (l *ZapLogger) Infof(msg string, args ...interface{}) {
	l.Logger.Sugar().Infof(msg, args...)
}

func (l *ZapLogger) Warn(msg string, fields ...interface{}) {
	l.Logger.Sugar().Warnw(msg, fields...)
}

func (l *ZapLogger) Warnf(msg string, args ...interface{}) {
	l.Logger.Sugar().Warnf(msg, args...)
}

func (l *ZapLogger) Error(msg string, fields ...interface{}) {
	l.Logger.Sugar().Errorw(msg, fields...)
}

func (l *ZapLogger) Errorf(msg string, fields ...interface{}) {
	l.Logger.Sugar().Errorf(msg, fields...)
}

func (l *ZapLogger) Fatal(msg string, fields ...interface{}) {
	l.Logger.Sugar().Fatalw(msg, fields...)
}

func (l *ZapLogger) Fatalf(msg string, fields ...interface{}) {
	l.Logger.Sugar().Fatalf(msg, fields...)
}

func (l *ZapLogger) Printf(msg string, args ...interface{}) {
	l.Logger.Sugar().Infof(msg, args...)
}

func (l *ZapLogger) Logf(msg string, args ...interface{}) {
	l.Logger.Sugar().Infof(msg, args...)
}

func DefaultLogger() Logger {
	logger, _ := NewZapLogger(WithLevel(zap.InfoLevel), WithEncodeTime("timestamp", zapcore.ISO8601TimeEncoder))
	return logger
}

func DefaultPrintfLogger() Logger {
	return &ZapLogger{Logger: zap.NewExample(), loggerConfig: zap.NewProductionConfig()}
}

func NewLogger(loggerType string, loggerLevel string) (Logger, error) {
	var err error

	switch loggerType {
	case "zap":
		var zapLevel zapcore.Level
		zapLevel, err = zapcore.ParseLevel(loggerLevel)
		if err != nil {
			return nil, fmt.Errorf("failed to parse logger level: %v", err)
		}
		return NewZapLogger(WithLevel(zapLevel), WithEncodeTime("timestamp", zapcore.ISO8601TimeEncoder))
	default:
		return nil, fmt.Errorf("unsupported logger type: %s", loggerType)
	}
}
