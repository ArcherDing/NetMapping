package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type NetMap struct {
	Protocol string
	From     string
	To       string
	Listen   string
	LogFile  string
}

type SvrConfig struct {
	configfile string
	Protocol   string
	Listen     string
	LogFile    string
	NetMapList []NetMap
}

func NewSvrConfig(configfile string) *SvrConfig {
	return &SvrConfig{
		configfile: configfile,
	}
}

func (self *SvrConfig) LoadConfig() error {
	file, err := os.Open(self.configfile)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(&self)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}

func (self *SvrConfig) DumpConfig() {
	fmt.Printf("%v", self)
}
