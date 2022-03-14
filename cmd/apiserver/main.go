package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"goCleanArch/internal/app/apiserver"
	"log"
)

var (
	configpath string
)

func init() {
	flag.StringVar(&configpath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	conf := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configpath, &conf); err != nil {
		log.Fatal("can't parse toml file ", err)
	}

	if err := apiserver.Start(conf); err != nil {
		log.Fatal("cant Start server ", err)
	}
}
