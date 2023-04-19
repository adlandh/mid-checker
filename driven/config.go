package driven

import (
	"github.com/caarlos0/env/v7"
)

type RetryerConfug struct {
	Timeout             int `env:"TIMEOUT" envDefault:"60"`
	TLSHandshakeTimeout int `env:"TLS_TIMEOUT" envDefault:"5"`
}

type Config struct {
	PassportFormID   string        `env:"ID,required"`
	WebhookURL       string        `env:"WEBHOOK_URL,required"`
	PeriodOfChecking int           `env:"PERIOD" envDefault:"3600"`
	Retryer          RetryerConfug `envPrefix:"RETRYER"`
	HttpPort         string        `env:"PORT" envDefault:"80"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg, env.Options{
		RequiredIfNoDef: false,
	})
	return cfg, err
}
