package configs

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
	DB MySQL
}

// MySQL represents configuration of MySQL database.
type MySQL struct {
	Username string `default:"root" envconfig:"MYSQL_USER"`
	Password string `default:"test" envconfig:"MYSQL_PASS"`
	Host     string `default:"127.0.0.1" envconfig:"MYSQL_HOST"`
	Port     int    `default:"3306" envconfig:"MYSQL_PORT"`
	Database string `default:"test" envconfig:"MYSQL_DB"`
}

// root:test@tcp(localhost:3306)/test

// ConnectionString returns connection string of MySQL database.
func (c *MySQL) ConnectionString() string {
	format := "%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8"
	return fmt.Sprintf(format, c.Username, c.Password, c.Host, c.Port, c.Database) + "&loc=Asia%2FHo_Chi_Minh"
}

// AppConfig is addition config for gami-service
var AppConfig App

// New returns a new instance of playground configuration.
func New() (*App, error) {
	if err := envconfig.Process(Prefix, &AppConfig); err != nil {
		return nil, err
	}
	return &AppConfig, nil
}
