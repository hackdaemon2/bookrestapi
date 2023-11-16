package models

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type appConfig struct {
	DatabaseUsername              string `yaml:"db_user"`
	DatabasePassword              string `yaml:"db_password"`
	DatabasePort                  string `yaml:"db_port"`
	DatabaseHost                  string `yaml:"db_host"`
	DatabaseName                  string `yaml:"db_name"`
	DatabaseMaximumOpenConnection int    `yaml:"db_max_open"`
	DatabaseMaximumIdleConnection int    `yaml:"db_max_idle_conn"`
	DatabaseMaximumIdleTime       int    `yaml:"db_max_idle_time"`
	DatabaseMaximumTime           int    `yaml:"db_max_time"`
	AppReadTimeout                int    `yaml:"app_read_timeout"`
	AppWriteTimeout               int    `yaml:"app_write_timeout"`
	AppJwtSecret                  string `yaml:"app_jwt_secret"`
	AppJwtTtl                     int    `yaml:"app_jwt_ttl"`
	AppJwtIssuer                  string `yaml:"app_jwt_issuer"`
	AppBearerToken                string `yaml:"bearer_token"`
}

// NewAppConfig creates a new AppConfig.
func NewAppConfig(envFile string) (IAppConfig, error) {
	var config appConfig

	// Load environment variables from the .env file
	err := cleanenv.ReadConfig(envFile, &config)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return &config, nil
}

func (config *appConfig) Password() string {
	return config.DatabasePassword
}

func (config *appConfig) Port() string {
	return config.DatabasePort
}

func (config *appConfig) Username() string {
	return config.DatabaseUsername
}

func (config *appConfig) Host() string {
	return config.DatabaseHost
}

func (config *appConfig) Name() string {
	return config.DatabaseName
}

func (config *appConfig) JwtSecret() string {
	return config.AppJwtSecret
}

func (config *appConfig) MaximumOpenConnection() int {
	return config.DatabaseMaximumOpenConnection
}

func (config *appConfig) MaximumIdleConnection() int {
	return config.DatabaseMaximumIdleConnection
}

func (config *appConfig) MaximumIdleTime() int {
	return config.DatabaseMaximumIdleTime
}

func (config *appConfig) MaximumTime() int {
	return config.DatabaseMaximumTime
}

func (config *appConfig) Issuer() string {
	return config.AppJwtIssuer
}

func (config *appConfig) ReadTimeout() int {
	return config.AppReadTimeout
}

func (config *appConfig) WriteTimeout() int {
	return config.AppWriteTimeout
}

func (config *appConfig) JwtTtl() int {
	return config.AppJwtTtl
}

func (config *appConfig) JwtIssuer() string {
	return config.AppJwtIssuer
}

func (config *appConfig) BearerToken() string {
	return config.AppBearerToken
}
