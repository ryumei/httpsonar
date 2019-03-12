package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type TargetConfig struct {
	URL          string
	ExpectedCode int
	Method       string
}

type TargetsConfig struct {
	Target []TargetConfig
}

// Config stores settings for application
type Config struct {
	Target []TargetConfig
	Log    LogConfig
}

func main() {
	confNamePtr := flag.String("conf", "httpsonar.toml", "config file")
	flag.Parse()

	log.Println("[INFO] Load config", *confNamePtr)
	var config Config
	_, err := toml.DecodeFile(*confNamePtr, &config)
	if err != nil {
		log.Println("[ERROR] Failed to load config file")
	}

	ConfigLogging(config.Log)

	for idx, target := range config.Target {
		result, err := ping(target.URL, target.ExpectedCode)
		fmt.Printf("entry:%d result:%v error:%v\n", idx, result, err)
	}
}
