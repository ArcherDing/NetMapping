package main

import (
	"flag"
	"log"
	//"net/http"
)

func main() {
	cfgfile := flag.String("conf", "config.json", "input a config file")
	from := flag.String("from", "", "input address of client")
	to := flag.String("-to", "", "input address of server")
	listen := flag.String("listen", "", "input address of nm-svr")
	flag.Usage = func(){
		
	}
	flag.Parse()
	cfg := NewCliConfig(*cfgfile)
	err := cfg.LoadConfig()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	//cfg.DumpConfig()
	client := NewClient(cfg)
	log.Println(flag.NArg())
	if flag.NArg() < 1{
		log.Println("ERRRR")
		return
	}
	switch flag.Arg(1) {
	case "list":
		client.List()
	case "add":
		netmap:= &NetMap{From:*from,To:*to,Listen:*listen}
		client.Add(netmap)
	case "update":
		client.Update()
	case "delete":
		client.Delete()
	case "remote":
		client.Remote()
	case "start":
		client.Start()
	case "stop":
		client.Stop()
	case "restart":
		client.Restart()
	case "pause":
		client.Pause()
	default:
		client.List()
	}
}
