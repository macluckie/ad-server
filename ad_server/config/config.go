package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PortGrpc           int64
	UrlRedis           string
	UrlTrackingService string
}

func NewConfig() (*Config, error) {
	config := Config{}
	port, err := getPortServer()
	if err != nil {
		return nil, err
	}
	config.PortGrpc = *port
	config.UrlRedis = *getRedisUrl()
	config.UrlTrackingService = *getUrlTrackingService()
	return &config, nil
}

func getPortServer() (*int64, error) {
	var port int64
	if os.Getenv("GRPC_PORT") != "" {
		portParsed, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
		if err != nil {
			return nil, fmt.Errorf("failed to parse GRPC_PORT error: %w", err)
		}
		port = int64(portParsed)
		return &port, nil
	}
	port = 50051
	return &port, nil
}

func getRedisUrl() *string {
	var url string
	if os.Getenv("REDIS_URL") != "" {
		url = os.Getenv("REDIS_URL")
		return &url
	}
	url = "redis://redis-data:6379/0?protocol=3"
	return &url
}

func getUrlTrackingService() *string {
	var url string
	if os.Getenv("TRACKING_URL") != "" {
		url = os.Getenv("TRACKING_URL")
		return &url
	}
	url = "tracking-service:50052"
	return &url
}
