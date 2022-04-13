package main

import (
	"github.com/sirupsen/logrus"
	"golang-ture"
	"golang-ture/internal/handler"
	"golang-ture/internal/repositories"
	"golang-ture/internal/services"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	storage := repositories.NewStorage()
	repository := repositories.NewRepository(storage)
	service := services.NewService(repository)
	handlers := handler.NewHandler(service)

	srv := new(golang_ture.Server)
	if err := srv.Run("8787", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error detected while running http server: %s", err.Error())
	}
}
