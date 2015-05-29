package main

type Client interface {
	List()
	Add(*NetMap)
	Update()
	Delete()
	Remote()
	Start()
	Stop()
	Restart()
	Pause()
}

type NetMap struct {
	Protocol string
	From     string
	To       string
	Listen   string
	LogFile  string
}


func NewClient(cfg *CliConfig) (client Client) {
	switch cfg.Protocol {
	case "Restful":
		client = &RestfulClient{ cfg : cfg }
	default:
		client = &RestfulClient{ cfg : cfg }
	}
	return
}