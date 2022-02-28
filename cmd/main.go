package main

import (
	yamlconfig "brbchat/plugins/config"
	"brbchat/plugins/obs"
	"brbchat/plugins/systray"
	"brbchat/plugins/twitchchat"
	"flag"
	"fmt"
)

var (
	filename = flag.String("config", "config.yaml", "filename of the config file")
)

func main() {
	flag.Parse()

	conf := yamlconfig.New(*filename)
	config, err := conf.Load()
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	st := systray.New()

	obs, err := obs.New(config, st)
	if err != nil {

	}
	twitchchat.New(config, obs, st).Listen()
}
