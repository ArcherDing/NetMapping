package main

import (
	"flag"
	"log"
	//"net/http"
)

var (
	cfgfile = flag.String("c", "config.json", "input a config file")
	logfile = flag.String("l", "nw-cli.log", "input a log file")
	cfg     *CliConfig
)

func main() {
	flag.Parse()
	cfg = NewCliConfig(*cfgfile)
	err := cfg.LoadConfig()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	//cfg.DumpConfig()
	switch cfg.Protocol {
	case "Restful":
		RunRestfulClient()
	default:
		RunRestfulClient()
	}
}
