package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type monitors struct {
	MonitorName string
	MonitorPath string
	MonitorPass string
	MonitorFail string
}

type configuration struct {
	PollInterval time.Duration
	ServerName   string
	Monitors     []monitors
}

var (
	conf configuration
)

func init() {

	//Load Default Config Values
	conf.PollInterval = 15 * time.Second
	conf.ServerName, _ = os.Hostname()

	//Load Configuration File
	dat, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Println("Failed to load conf.json with error: \n\n--", err, "\n ")
	} else {
		err = json.Unmarshal(dat, &conf)
	}
}
