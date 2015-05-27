package main

import (
	"flag"
	"log"
)

var (
	cfgfile = flag.String("c", "config.json", "input a config file")
	logfile = flag.String("l", "nw-svr.log", "input a log file")
	cfg     *SvrConfig
)

func main() {
	flag.Parse()
	cfg = NewSvrConfig(*cfgfile)
	err := cfg.LoadConfig()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	//cfg.DumpConfig()

	switch cfg.Protocol {
	case "Restful":
		RunRestfulServer()
	default:
		RunRestfulServer()
	}
}
