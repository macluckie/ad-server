package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"test/grpc/config"
	"test/grpc/domain"

	pb "test/grpc/proto/tracking"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service interface {
	GetCount(ctx context.Context, id *domain.IdAd) (*domain.CountAd, error)
	IncrementCount(ctx context.Context, id *domain.IdAd) error
}

type TrackingServer struct {
	trackingService Service
	config          *config.Config
	pb.UnimplementedTrackingServerServer
}

func NewTrackingServer(config *config.Config, trackingService Service) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.PortGrpc))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	trackServer := TrackingServer{
		config:          config,
		trackingService: trackingService,
	}
	pb.RegisterTrackingServerServer(grpcServer, &trackServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (t *TrackingServer) GetCountAd(ctx context.Context, in *pb.IdAdTracked) (*pb.Count, error) {
	id := domain.IdAd(in.Id)
	count, err := t.trackingService.GetCount(ctx, &id)
	if err != nil {
		return nil, err
	}
	response := pb.Count{
		Count: int64(*count),
	}
	return &response, nil
}

func (t *TrackingServer) IncrementCount(ctx context.Context, in *pb.IdAdTracked) (*emptypb.Empty, error) {
	err := t.trackingService.IncrementCount(ctx, (*domain.IdAd)(&in.Id))
	if err != nil {
		return nil, err
	}
	response := emptypb.Empty{}
	return &response, nil
}
