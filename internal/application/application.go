package application

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"personal_knowledge_tracker/config"
	"personal_knowledge_tracker/internal/interfaces"
	"personal_knowledge_tracker/internal/usecases"
	"personal_knowledge_tracker/pkg/database/mongodb"
	"personal_knowledge_tracker/pkg/logger/zap"
	"syscall"

	"github.com/go-logr/logr"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	cfg        *config.Config
	httpServer interfaces.Server
	logger     logr.Logger
	*mongodb.MongoDB

	shutdown chan os.Signal
}

func NewWithContext(ctx context.Context, cfg *config.Config) (*Application, error) {
	mongodb, err := mongodb.NewMongoDB(cfg.Mongo.DSN(), cfg.Mongo.DB)
	if err != nil {
		return nil, err
	}

	app := &Application{
		ctx:      ctx,
		cfg:      cfg,
		MongoDB:  mongodb,
		shutdown: make(chan os.Signal, 1),
	}

	signal.Notify(app.shutdown, os.Interrupt, syscall.SIGTERM)

	return app, nil
}

func (a *Application) Run() {
	a.logger.Info("Application started...")

	<-a.shutdown
	a.Stop()
}

func (a *Application) Stop() {
	a.cancelFunc()
	close(a.shutdown)

	a.logger.Info("Application stoped")
}

func (a *Application) InitUsecases() interfaces.Usecases {
	ctx, cancelFunc := context.WithCancel(a.ctx)
	a.cancelFunc = cancelFunc

	deps := usecases.Dependencies{
		Ctx:    ctx,
		Logger: a.logger,
	}

	return usecases.New(deps)
}

func (a *Application) InitLogger() {
	a.logger = zap.New(
		zap.Level(a.cfg.Level),
		zap.UseDevMode(true),
		zap.TimeEncoder(zapcore.ISO8601TimeEncoder),
		zap.ConsoleEncoder(
			func(ec *zapcore.EncoderConfig) { ec.EncodeLevel = zapcore.CapitalColorLevelEncoder },
			func(ec *zapcore.EncoderConfig) {
				ec.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
			},
			func(ec *zapcore.EncoderConfig) {
				ec.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
					encoder.AppendString(filepath.Base(caller.FullPath()))
				}
			},
		),
	)
}

func (a *Application) GetLogger() logr.Logger {
	return a.logger
}

func (a *Application) RegisterHTTPServer(server interfaces.Server) {
	a.httpServer = server
}
