package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var RobotCfg struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	MgodbAddr 	string
	GameMgoConnNum  int
	LoginMgoConnNum  int
	ProfilePath string
}

func init() {
	data, err := ioutil.ReadFile("common.json")
	if err != nil {
		log.Printf("%v", err)
	}
	err = json.Unmarshal(data, &RobotCfg)
	if err != nil {
		log.Printf("%v", err)
	}
}
