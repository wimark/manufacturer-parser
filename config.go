package main

import (
	log "github.com/wimark/liblog"

	"github.com/kelseyhightower/envconfig"
)

//Config конфигаруционные данные приложения
type Config struct {
	PORT    string `envconfig:"PORT" default:":7755"`
	DBURL   string `envconfig:"DBURL" required:"true"`
	DATAURL string `envconfig:"DATAURL" default:"https://linuxnet.ca/ieee/oui/nmap-mac-prefixes"`
}

//Init считывает конфиги с env var
func (c *Config) Init() {
	err := envconfig.Process("", c)
	if err != nil || len(c.DBURL) == 0 {
		panic(err)
	}
	log.Info("config: %+v", c)
}
