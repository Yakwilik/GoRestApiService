package main

import (
	"log"

	todo "github.com/Yakwilik/GoRestApiServiceToDo"

	handler "github.com/Yakwilik/GoRestApiServiceToDo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("произошла ошибка во время запуска http-сервера: %s", err.Error())
	}
}
