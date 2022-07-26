package main

import (
	"log"

	todo "github.com/Yakwilik/GoRestApiServiceToDo"
	"github.com/spf13/viper"

	handler "github.com/Yakwilik/GoRestApiServiceToDo/pkg/handler"
	"github.com/Yakwilik/GoRestApiServiceToDo/pkg/repository"
	"github.com/Yakwilik/GoRestApiServiceToDo/pkg/service"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("Ошибка в прочтении файла конфигурации: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("произошла ошибка во время запуска http-сервера: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
