package models

type Config struct {
	Port string
	DB   Postgres `envconfig:"POSTGRES"`
}

type Postgres struct {
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Host     string `envconfig:"HOST" required:"true"` // localhost
	Port     string `envconfig:"PORT" required:"true"` //:5432
	Database string `envconfig:"DATABASE" required:"true"`
	SSLMode  string `envconfig:"SSL_MODE" default:"disable"`
}
