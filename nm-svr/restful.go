package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RunRestfulServer() {
	m := mux.NewRouter()
	m.HandleFunc("/config", GetConfig).Methods("GET")
	m.HandleFunc("/netmaps", GetNetMapList).Methods("GET")
	m.HandleFunc("/netmaps", AddNetMap).Methods("POST")
	m.HandleFunc("/netmaps/{id:[0-9]+}", UpdateNetMap).Methods("PUT")
	m.HandleFunc("/netmaps/{id:[0-9]+}", DelNetMap).Methods("DELETE")

	log.Println("nm-svr is start on", cfg.Listen)
	err := http.ListenAndServe(cfg.Listen, m)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	_, err = w.Write(bytes)
	if err != nil {
		log.Println("error", "Failed to write: %v", err.Error())
		log.Println("info", "Message was: %v", string(bytes))
	}
}

func GetNetMapList(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(cfg.NetMapList)
	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	_, err = w.Write(bytes)
	if err != nil {
		log.Println("error", "Failed to write: %v", err.Error())
		log.Println("info", "Message was: %v", string(bytes))
	}
}

func AddNetMap(w http.ResponseWriter, r *http.Request) {

}

func UpdateNetMap(w http.ResponseWriter, r *http.Request) {

}

func DelNetMap(w http.ResponseWriter, r *http.Request) {

}
