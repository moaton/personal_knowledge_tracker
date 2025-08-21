package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap/zapcore"
)

type (
	Config struct {
		Log  `yaml:"logger"`
		HTTP `yaml:"http"`
		CORS `yaml:"cors"`
	}

	Log struct {
		Level zapcore.Level `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}
	HTTP struct {
		Server struct {
			Port         string `env-required:"true" yaml:"http_server_port" env:"HTTP_SERVER_PORT"`
			WriteTimeout int    `env-required:"true" yaml:"http_server_write_timeout" env:"HTTP_SERVER_WRITE_TIMEOUT" env-default:"30"`
		} `yaml:"server"`
	}
	CORS struct {
		AllowMethods          string `env-required:"true" yaml:"cors_allow_methods" env:"CORS_ALLOW_METHODS"`
		AllowOrigin           string `env-required:"true" yaml:"cors_allow_origin" env:"CORS_ALLOW_ORIGIN"`
		AllowCredentials      string `env-required:"true" yaml:"cors_allow_credentials" env:"CORS_ALLOW_CREDENTIALS"`
		AllowHeaders          string `env-required:"true" yaml:"cors_allow_headers" env:"CORS_ALLOW_HEADERS"`
		XContentTypeOptions   string `env-required:"true" yaml:"cors_x_content_type_options" env:"CORS_X_CONTENT_TYPE_OPTIONS"`
		XFrameOptions         string `env-required:"true" yaml:"cors_x_frame_options" env:"CORS_X_FRAME_OPTIONS"`
		ContentSecurityPolicy string `env-required:"true" yaml:"cors_content_security_policy" env:"CORS_CONTENT_SECURITY_POLICY"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("failed to load env from file: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
