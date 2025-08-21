package zap

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New -.
func New(opts ...Opts) logr.Logger {
	return zapr.NewLogger(newZap(opts...))
}

func newZap(opts ...Opts) *zap.Logger {
	o := &Options{}

	for _, opt := range opts {
		opt(o)
	}

	o.addDefaults()

	sink := zapcore.AddSync(o.DestWriter)

	o.ZapOpts = append(o.ZapOpts, zap.ErrorOutput(sink), zap.AddCaller())
	zl := zap.New(zapcore.NewCore(o.Encoder, sink, o.Level))
	zl = zl.WithOptions(o.ZapOpts...)

	return zl
}
