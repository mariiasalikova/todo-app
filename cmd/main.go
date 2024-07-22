package main

import (
	"github.com/mariiasalikova/todo-app"
	"github.com/mariiasalikova/todo-app/pkg/handler"
	"github.com/mariiasalikova/todo-app/pkg/repository"
	"github.com/mariiasalikova/todo-app/pkg/service"
	"github.com/github.com/spf13/viper"
	"log")

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}


	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running the server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")

	return viper.ReadInConfig()

}