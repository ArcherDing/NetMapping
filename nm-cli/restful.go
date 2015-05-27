package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
)

type NetMap struct {
	Protocol string
	From     string
	To       string
	Listen   string
	LogFile  string
}

var netMapList []NetMap

func RunRestfulClient() {
	resp, err := http.Get("http://"+cfg.Server+"/netmaps")
	if err != nil {
		log.Println("Server is not started, please run `nm-svr` on server!")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	
	err = json.Unmarshal(body,&netMapList)
	if err != nil {
		log.Println(err.Error())
	}
	
	fmt.Printf("%2s\t%-10s\t%-20s\t%-20s\t%-20s\n","No","Protocol","From","To","Listen")
	for i,v := range netMapList {
		fmt.Printf("%02d\t%-10s\t%-20s\t%-20s\t%-20s\n",i,v.Protocol,v.From,v.To,v.Listen)
	}
}