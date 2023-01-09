package main

import (
	"rahmathidayat1203/backend-portofolio-v1/config"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Infof(".env not loaded")
	}
	log.Infof("read .env file")

	config := config.NewConfig()
	server := InitServer(config)

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()
		server.Run()
	}()
	wg.Wait()
}
