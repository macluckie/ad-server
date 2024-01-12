package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"test/grpc/config"
	"test/grpc/domain"
	pb "test/grpc/proto/ad"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service interface {
	Create(ctx context.Context, ad *domain.Ad) error
	Get(ctx context.Context, id *domain.IdAd) (*domain.Ad, error)
	ServeAd(ctx context.Context, id *domain.IdAd) (*domain.AdServe, error)
}

type AdServer struct {
	adService Service
	config    *config.Config
	validate  *validator.Validate
	pb.UnimplementedAdServerServer
}

func NewAdServer(config *config.Config, validator *validator.Validate, adService Service) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.PortGrpc))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	adServer := AdServer{
		config:    config,
		validate:  validator,
		adService: adService,
	}
	pb.RegisterAdServerServer(grpcServer, &adServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *AdServer) CreateAd(ctx context.Context, in *pb.AdRequest) (*emptypb.Empty, error) {
	ad := domain.Ad{
		Id:          in.Id,
		Title:       in.Title,
		URL:         in.Url,
		Description: in.Description,
	}
	err := s.checkStruct(&ad)
	if err != nil {
		return nil, err
	}
	err = s.adService.Create(ctx, &ad)
	if err != nil {
		return nil, err
	}
	response := emptypb.Empty{}
	return &response, nil
}

func (s *AdServer) checkStruct(ad *domain.Ad) error {
	err := s.validate.Struct(ad)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Validation error in field '%s': %s\n", err.Field(), err.Tag())
		}
		return fmt.Errorf("failed to validate Ad")
	}
	return nil
}

func (s *AdServer) GetAd(ctx context.Context, in *pb.IdAd) (*pb.AdResponse, error) {
	id := domain.IdAd(in.Id)
	ad, err := s.adService.Get(ctx, &id)
	if err != nil {
		return nil, err
	}
	response := pb.AdResponse{
		Id:          ad.Id,
		Title:       ad.Title,
		Description: ad.Description,
		Url:         ad.URL,
	}
	return &response, nil
}

func (s *AdServer) ServeAd(ctx context.Context, in *pb.IdAd) (*pb.Ad, error) {
	id := domain.IdAd(in.Id)
	adServe, err := s.adService.ServeAd(ctx, &id)
	if err != nil {
		return nil, err
	}
	response := pb.Ad{
		Url:      adServe.Url,
		Tracking: adServe.TrackImpression,
	}
	return &response, nil
}
