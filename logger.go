package main

import (
	"github.com/hashicorp/logutils"
	"log"
	"os"
	"path"
)

// LogConfig stores log settings
type LogConfig struct {
	LogFile string
	Level   string
}

func prepareLogDir(logpath string) error {
	parent := path.Dir(logpath)
	if err := os.MkdirAll(parent, os.ModePerm); err != nil {
		log.Printf("[ERROR] %v", err)
		return err
	}
	return nil
}

func openLogFile(logPath string) *os.File {
	prepareLogDir(logPath)
	if logPath == "" {
		log.Printf("[WARN] [Server] ServerLog is undefined. Log to STDERR.")
		return os.Stderr
	}
	logWriter, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("[WARN] %v Log to STDERR.", err)
		return os.Stderr
	}
	return logWriter
}

// ConfigLogging set
func ConfigLogging(conf LogConfig) {
	logWriter := openLogFile(conf.LogFile)

	logLevel := conf.Level
	if logLevel == "" {
		logLevel = "INFO"
	}

	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(logLevel),
		Writer:   logWriter,
	}
	log.SetOutput(filter)
	logFlags := log.LstdFlags | log.Lmicroseconds | log.LUTC
	if filter.MinLevel == "DEBUG" {
		logFlags |= log.Lshortfile
	}
	log.SetFlags(logFlags)
}
