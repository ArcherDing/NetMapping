package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
)

type RestfulClient struct {
	cfg *CliConfig
}

func (self *RestfulClient)List() {
	resp, err := http.Get("http://"+self.cfg.Server+"/netmaps")
	if err != nil {
		log.Println("Server is not started, please run `nm-svr` on server!")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	
	var netMapList []NetMap
	err = json.Unmarshal(body,&netMapList)
	if err != nil {
		log.Println(err.Error())
	}
	
	fmt.Printf("%2s\t%-10s\t%-20s\t%-20s\t%-20s\n","No","Protocol","From","To","Listen")
	for i,v := range netMapList {
		fmt.Printf("%02d\t%-10s\t%-20s\t%-20s\t%-20s\n",i,v.Protocol,v.From,v.To,v.Listen)
	}
}

func (self *RestfulClient)Add(nm *NetMap) {
}


func (self *RestfulClient)Update() {
}
func (self *RestfulClient)Delete() {
}
func (self *RestfulClient)Remote() {
}
func (self *RestfulClient)Start() {
}
func (self *RestfulClient)Stop() {
}
func (self *RestfulClient)Restart() {
}
func (self *RestfulClient)Pause() {
}