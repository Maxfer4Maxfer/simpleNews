package main

import (
	log "github.com/sirupsen/logrus"

	"simpleNews/pkg/config"
	server "simpleNews/pkg/server/nats"
	"simpleNews/pkg/service"
	storage "simpleNews/pkg/storage/gorm"
)

func main() {

	cfg := config.GetConfig()

	// setup logger
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{})
	// log.SetReportCaller(true)

	storage := storage.New(cfg)
	defer storage.Close()

	service := service.New(storage)

	server := server.New(cfg, service)

	server.Run()
}
