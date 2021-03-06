package main

import (
	"log"
	"os"
	"path/filepath"
)

type FileLogger struct {
	*log.Logger
	*os.File
}

func NewFileLogger(path string) (*FileLogger, error) {
	logFile := &FileLogger{}
	err := os.MkdirAll(filepath.Dir(hkPath()), 0777)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	logFile.File = file
	logFile.Logger = log.New(file, "", log.LstdFlags)
	return logFile, nil
}
