package main

import (
	"log"

	todo "github.com/Yakwilik/GoRestApiServiceToDo"

	handler "github.com/Yakwilik/GoRestApiServiceToDo/pkg/handler"
	"github.com/Yakwilik/GoRestApiServiceToDo/pkg/repository"
	"github.com/Yakwilik/GoRestApiServiceToDo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("произошла ошибка во время запуска http-сервера: %s", err.Error())
	}
}
