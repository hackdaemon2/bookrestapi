package models

type IAppConfig interface {
	Username() string
	Password() string
	Port() string
	Host() string
	Name() string
	MaximumOpenConnection() int
	MaximumIdleConnection() int
	MaximumIdleTime() int
	MaximumTime() int
	ReadTimeout() int
	WriteTimeout() int
	JwtSecret() string
	JwtTtl() int
	JwtIssuer() string
	BearerToken() string
}
