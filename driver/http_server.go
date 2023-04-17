package driver

import (
	"context"
	"net/http"

	"github.com/adlandh/mid-checker/driven"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewHttpServer(lc fx.Lifecycle, cfg *driven.Config) (*echo.Echo, error) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			go func() {
				err = e.Start(":" + cfg.HttpPort)
			}()
			return err
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
	return e, nil
}
