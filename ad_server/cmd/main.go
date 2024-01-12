package main

import (
	"test/grpc/api"
	"test/grpc/config"
	"test/grpc/domain"
	"test/grpc/repo"

	"github.com/go-playground/validator/v10"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	tracking, err := repo.NewTrackingService(config)
	if err != nil {
		panic(err)
	}
	db, err := repo.NewRedisDb(config, tracking)
	if err != nil {
		panic(err)
	}
	service := domain.NewService(db)
	validator := validator.New()
	err = api.NewAdServer(config, validator, service)
	if err != nil {
		panic(err)
	}
}
