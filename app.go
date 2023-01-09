package main

import (
	"fmt"
	"log"
	"net/http"
	"rahmathidayat1203/backend-portofolio-v1/config"
	"rahmathidayat1203/backend-portofolio-v1/delivery"
	"rahmathidayat1203/backend-portofolio-v1/repository"
	"rahmathidayat1203/backend-portofolio-v1/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	httpServer *echo.Echo
	cfg        config.Config
}
type Server interface {
	Run()
}

func InitServer(cfg config.Config) Server {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &server{
		httpServer: e,
		cfg:        cfg,
	}
}
func (s *server) Run() {
	s.httpServer.GET("", func(e echo.Context) error {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"status":  "success",
			"message": "Backend Portof !" + s.cfg.ServiceName() + " " + s.cfg.ServiceEnvironment(),
		})
	})

	messageRepo := repository.NewMessageRepository(s.cfg)
	messageUsecase := usecase.NewMessageUsecase(messageRepo)
	messageDelivery := delivery.NewMessageDelivery(messageUsecase)
	messageGroup := s.httpServer.Group("/message")
	messageDelivery.Mount(messageGroup)

	if err := s.httpServer.Start(fmt.Sprintf(":%d", s.cfg.ServicePort())); err != nil {
		log.Panic(err)
	}
}
