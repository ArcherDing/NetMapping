package main

import (
	"flag"
	"fmt"
)

var (
	cfgfile = flag.String("c","config.json","input a config file")
	logfile = flag.String("l","nw-svr.log","input a log file")
)

func main() {
	flag.Parse()
	cfg := NewSvrConfig(*cfgfile)
	err := cfg.LoadConfig()
	if err != nil {
		log.Error(err.Error())
		return
	}

}
