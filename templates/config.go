package templates

// ConfigTemplate ...
const ConfigTemplate = `package configs

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Package constants definition.
const (
	Prefix = ""
)

// App represents all configuration of qr-service application.
type App struct {
	Host     string BACK_STICKdefault:"0.0.0.0" envconfig:"HOST"BACK_STICK
	Port     int    BACK_STICKdefault:"8080" envconfig:"PORT"BACK_STICK
	RunMode  string BACK_STICKdefault:"debug" envconfig:"RUN_MODE"BACK_STICK
	Env      string BACK_STICKdefault:"debug" envconfig:"ENV"BACK_STICK
	DB       MySQL
}

// MySQL represents configuration of MySQL database.
type MySQL struct {
	Username     string BACK_STICKdefault:"root" envconfig:"MYSQL_USER"BACK_STICK
	Password     string BACK_STICKdefault:"root" envconfig:"MYSQL_PASS"BACK_STICK
	Host         string BACK_STICKdefault:"127.0.0.1" envconfig:"MYSQL_HOST"BACK_STICK
	Port         int    BACK_STICKdefault:"3306" envconfig:"MYSQL_PORT"BACK_STICK
	Database     string BACK_STICKdefault:"name" envconfig:"MYSQL_DB"BACK_STICK
}

// ConnectionString returns connection string of MySQL database.
func (c *MySQL) ConnectionString() string {
	format := "%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8"
	return fmt.Sprintf(format, c.Username, c.Password, c.Host, c.Port, c.Database) + "&loc=Asia%2FHo_Chi_Minh"
}

// AppConfig is addition config for gami-service
var AppConfig App

// New returns a new instance of MODULE_NAME configuration.
func New() (*App, error) {
	if err := envconfig.Process(Prefix, &AppConfig); err != nil {
		return nil, err
	}
	return &AppConfig, nil
}

// AddressListener returns address listener of HTTP server.
func (c *App) AddressListener() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}
`