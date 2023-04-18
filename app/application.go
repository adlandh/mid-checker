package app

import (
	"context"
	"time"

	"github.com/adlandh/mid-checker/domain"
	"github.com/adlandh/mid-checker/driven"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Application struct {
	infoFetcher       domain.PassportInfoFetcher
	evenSender        domain.PassportChangedEventSender
	savedPassportInfo domain.PassportInfo
	logger            *zap.Logger
	period            time.Duration
}

func NewApplication(lc fx.Lifecycle, passportInfoFetcher domain.PassportInfoFetcher,
	eventSender domain.PassportChangedEventSender, logger *zap.Logger, cfg *driven.Config) *Application {
	app := &Application{
		infoFetcher: passportInfoFetcher,
		evenSender:  eventSender,
		logger:      logger,
		period:      time.Duration(cfg.PeriodOfChecking) * time.Second,
	}
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		go app.CheckStatusPeriodically()
		return nil
	}})

	return app
}

func (a *Application) CheckStatusPeriodically() {
	for {
		a.checkStatus()
		time.Sleep(a.period)
	}
}

func (a *Application) checkStatus() {
	a.logger.Info("start checking passport status")
	defer a.logger.Info("finish checking passport status")
	info, err := a.infoFetcher.GetPassportStatus()
	if err != nil {
		a.logger.Error("error checking passport status", zap.Error(err))
		return
	}

	if info == nil {
		a.logger.Error("empty passport status returned")
		return
	}

	if a.savedPassportInfo.Uid != "" &&
		(info.InternalStatus.Percent != a.savedPassportInfo.InternalStatus.Percent ||
			info.InternalStatus.Name != a.savedPassportInfo.InternalStatus.Name) {
		a.logger.Info("password changed, sending event")
		err := a.evenSender.SendChangedStatus(info)
		if err != nil {
			a.logger.Error("error sending event", zap.Error(err))
		}
	}

	a.savedPassportInfo = *info
}
