package main

import (
	"github.com/RobinSoGood/bookly/internal/config"
	"github.com/RobinSoGood/bookly/internal/logger"
	"github.com/RobinSoGood/bookly/internal/server"
	"github.com/RobinSoGood/bookly/internal/service"
	"github.com/RobinSoGood/bookly/internal/storage"
)

func main() {
	cfg := config.ReadConfig()
	log := logger.Get(cfg.Debug)

	stor := storage.New()
	userService := service.NewUserService(stor)
	bookService := service.NewBookService(stor)
	serve := server.New(cfg, userService, bookService)
	if err := serve.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
