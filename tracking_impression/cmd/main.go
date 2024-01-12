package main

import (
	"test/grpc/api"
	"test/grpc/config"
	"test/grpc/domain"
	"test/grpc/repo"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	db, err := repo.NewRedisDb(config)
	if err != nil {
		panic(err)
	}
	service := domain.NewService(db)
	err = api.NewTrackingServer(config, service)
	if err != nil {
		panic(err)
	}
}
