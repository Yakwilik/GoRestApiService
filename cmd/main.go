package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

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

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка при чтении переменной окружения: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Ошибка в инициализации базы данных: %s", err.Error())
	}
	repos := repository.NewRepository(db)
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
