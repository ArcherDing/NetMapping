package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type CliConfig struct {
	configfile string
	Protocol   string
	Server     string
	LogFile    string
}

func NewCliConfig(configfile string) *CliConfig {
	return &CliConfig{
		configfile: configfile,
	}
}

func (self *CliConfig) LoadConfig() error {
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

func (self *CliConfig) DumpConfig() {
	fmt.Printf("%v", self)
}
