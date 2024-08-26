package main

import (
	"kp/config"
	"kp/internal/api/handler"
	apirepository "kp/internal/api/repository"
	apiusecase "kp/internal/api/usecase"
	"kp/internal/router"
	"kp/pkg/datasource"
	"kp/pkg/logger"
)

func init() {
	config.InitConfig()
	logger.InitLogger()
}

func main() {
	db := datasource.NewGorm()

	repo := apirepository.New(db)
	uc := apiusecase.New(repo)
	handler := handler.New(uc)

	router := router.NewRouter(handler)
	router.MapHandler()

	router.Run()
}
