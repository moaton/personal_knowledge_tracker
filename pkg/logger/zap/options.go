package zap

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	EncoderConfigOption func(*zapcore.EncoderConfig)
	NewEncoderFunc      func(...EncoderConfigOption) zapcore.Encoder
	Opts                func(*Options)
)

// Опции логгера
type Options struct {
	// Development указывает под какую среду настраивается логгер (разработка или продуктовая).
	Development bool
	// Level указывает уровень логирования.
	Level zapcore.LevelEnabler
	// DestWriter устанавливает, куда будет записываться лок.  По умолчанию os.Stderr.
	DestWriter io.Writer
	// Encoder устанавливает, как zap будет выводить лог.
	// Для среды разработки по умолчанию консоль, иначе JSON.
	Encoder zapcore.Encoder
	// EncoderConfigOptions позволяют модифицировать опции вывода лога. Не применяются, если Encoder уже установлен.
	EncoderConfigOptions []EncoderConfigOption
	// NewEncoder настраивает Encoder в зависимости от EncoderConfigOptions. Не используется, если Encoder уже установлен.
	NewEncoder NewEncoderFunc
	// TimeEncoder позволяет настраивать формат вывода времени. По умолчанию EpochTimeEncoder.
	TimeEncoder zapcore.TimeEncoder
	// StacktraceLevel определяет, начиная с какого уровня логирования выводить стэк вызовов.
	StacktraceLevel zapcore.LevelEnabler
	// ZapOpts позволяет напрямую передать опции в zap логгер.
	ZapOpts []zap.Option
}

// addDefaults - добавляет в Options значения по умолчанию
func (o *Options) addDefaults() {
	if o.DestWriter == nil {
		o.DestWriter = os.Stderr
	}

	if o.Development {
		if o.Level == nil {
			o.Level = newAtomicLevelAt(zapcore.DebugLevel)
		}

		if o.NewEncoder == nil {
			o.NewEncoder = newConsoleEncoder
		}

		if o.StacktraceLevel == nil {
			o.StacktraceLevel = newAtomicLevelAt(zapcore.ErrorLevel)
		}

		o.ZapOpts = append(o.ZapOpts, zap.Development())
	} else {
		if o.Level == nil {
			o.Level = newAtomicLevelAt(zapcore.InfoLevel)
		}

		if o.NewEncoder == nil {
			o.NewEncoder = newJSONEncoder
		}

		if o.StacktraceLevel == nil {
			o.StacktraceLevel = newAtomicLevelAt(zapcore.FatalLevel)
		}
	}

	if o.TimeEncoder == nil {
		o.TimeEncoder = zapcore.EpochTimeEncoder
	}

	f := func(ecfg *zapcore.EncoderConfig) {
		ecfg.EncodeTime = o.TimeEncoder
	}

	o.EncoderConfigOptions = append([]EncoderConfigOption{f}, o.EncoderConfigOptions...)

	if o.Encoder == nil {
		o.Encoder = o.NewEncoder(o.EncoderConfigOptions...)
	}

	o.ZapOpts = append(o.ZapOpts, zap.AddStacktrace(o.StacktraceLevel))
}

func newAtomicLevelAt(level zapcore.Level) *zap.AtomicLevel {
	lvl := zap.NewAtomicLevelAt(level)
	return &lvl
}

// ConsoleEncoder - настраивает логгер для записи в консоль.
func ConsoleEncoder(opts ...EncoderConfigOption) Opts {
	return func(o *Options) {
		o.Encoder = newConsoleEncoder(opts...)
	}
}

// JSONEncoder - настраивает логгер для записи в JSON.
func JSONEncoder(opts ...EncoderConfigOption) Opts {
	return func(o *Options) {
		o.Encoder = newJSONEncoder(opts...)
	}
}

func newConsoleEncoder(opts ...EncoderConfigOption) zapcore.Encoder {
	encoderCfg := zap.NewDevelopmentEncoderConfig()

	for _, opt := range opts {
		opt(&encoderCfg)
	}

	return zapcore.NewConsoleEncoder(encoderCfg)
}

func newJSONEncoder(opts ...EncoderConfigOption) zapcore.Encoder {
	encoderCfg := zap.NewProductionEncoderConfig()

	for _, opt := range opts {
		opt(&encoderCfg)
	}

	return zapcore.NewJSONEncoder(encoderCfg)
}

// UseDevMode -.
func UseDevMode(enabled bool) Opts {
	return func(o *Options) { o.Development = enabled }
}

// Level - устанавливает уровень логирования. Минимальный уровень есть Debug.
// Соотношение уровней logr и zap указаны здесь: https://pkg.go.dev/github.com/go-logr/zapr
func Level(level zapcore.Level) Opts {
	return func(o *Options) {
		if level.Enabled(zapcore.Level(-2)) {
			level = zapcore.DebugLevel
		}
		o.Level = level
	}
}

// WriteTo - устанавливает, куда будет записываться лог
func WriteTo(out io.Writer) Opts {
	return func(o *Options) {
		o.DestWriter = out
	}
}

// StacktraceLevel - устанавливает уровень логирования, начиная с которого выводить стэк вызовов.
func StacktraceLevel(stacktraceLevel zapcore.LevelEnabler) Opts {
	return func(o *Options) {
		o.StacktraceLevel = stacktraceLevel
	}
}

func TimeEncoder(timeEncoder zapcore.TimeEncoder) Opts {
	return func(o *Options) {
		o.TimeEncoder = timeEncoder
	}
}

// RawZapOpts - настраивает zap напрямую через zap.Option.
func RawZapOpts(zapOpts ...zap.Option) Opts {
	return func(o *Options) {
		o.ZapOpts = append(o.ZapOpts, zapOpts...)
	}
}
