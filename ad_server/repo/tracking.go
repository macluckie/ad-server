package repo

import (
	"fmt"
	"test/grpc/config"
	pb_tracking "test/grpc/proto/tracking"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Tracking struct {
	config *config.Config
	client pb_tracking.TrackingServerClient
}

func NewTrackingService(config *config.Config) (*Tracking, error) {
	conn, err := grpc.Dial(config.UrlTrackingService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to tracker-service, error: %w", err)
	}
	service := pb_tracking.NewTrackingServerClient(conn)
	trackingService := Tracking{
		config: config,
		client: service,
	}
	return &trackingService, nil
}
