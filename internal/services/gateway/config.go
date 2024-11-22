package gateway

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	AppEnv             string `env:"APP_ENV" env-default:"test"`
	Port               int    `env:"PORT" env-default:"8080"`
	LoyaltyService     string `env:"LOYALTY_SERVICE"`
	PaymentService     string `env:"PAYMENT_SERVICE"`
	ReservationService string `env:"RESERVATION_SERVICE"`
}

func NewConfig() (*Config, error) {
	localPath := "./configs/gateway.env"
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	if cfg.AppEnv != "test" {
		return nil, nil
	}

	err = cleanenv.ReadConfig(localPath, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
