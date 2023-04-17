package main

import (
	"github.com/adlandh/mid-checker/app"
	"github.com/adlandh/mid-checker/domain"
	"github.com/adlandh/mid-checker/driven"
	"github.com/adlandh/mid-checker/driver"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			driven.NewConfig,
			driven.NewHttpClient,
			driven.NewLogger,
			fx.Annotate(
				driven.NewMidServiceClient,
				fx.As(new(domain.PassportInfoFetcher)),
			),
			fx.Annotate(
				driven.NewWebhookEventSender,
				fx.As(new(domain.PassportChangedEventSender)),
			),
		),
		fx.Invoke(app.NewApplication, driver.NewHttpServer),
	).Run()
}
