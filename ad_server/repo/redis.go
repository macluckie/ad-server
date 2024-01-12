package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"test/grpc/config"
	"test/grpc/domain"
	"time"

	pb_tracking "test/grpc/proto/tracking"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	client   *redis.Client
	tracking *Tracking
}

func NewRedisDb(config *config.Config, tracking *Tracking) (*RedisDB, error) {
	opts, err := redis.ParseURL(config.UrlRedis)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url redis: %w", err)
	}
	db := RedisDB{
		client:   redis.NewClient(opts),
		tracking: tracking,
	}
	ctx := context.Background()
	_, err = db.client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis, error: %w", err)
	}
	return &db, nil
}

func (r *RedisDB) Create(ctx context.Context, ad *domain.Ad) error {
	data, err := json.Marshal(ad)
	if err != nil {
		return fmt.Errorf("failed to Marshall ad, error: %w", err)
	}
	err = r.client.Set(ctx, ad.Id, data, 120*time.Second).Err()
	if err != nil {
		return fmt.Errorf("failed to create Ad in database, error: %w", err)
	}
	return nil
}

func (r *RedisDB) Get(ctx context.Context, id *domain.IdAd) (*domain.Ad, error) {
	response, err := r.client.Get(ctx, string(*id)).Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to query Ad in database, error: %w", err)
	}
	ad := domain.Ad{}
	err = json.Unmarshal(response, &ad)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall response, error: %w", err)
	}
	return &ad, nil
}

func (r *RedisDB) ServeAd(ctx context.Context, id *domain.IdAd) (*domain.AdServe, error) {
	ad, err := r.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	adId := pb_tracking.IdAdTracked{
		Id: ad.Id,
	}
	_, err = r.tracking.client.IncrementCount(ctx, &adId)
	if err != nil {
		return nil, err
	}
	countTrackingImpression, err := r.tracking.client.GetCountAd(ctx, &adId)
	if err != nil {
		return nil, err
	}
	response := domain.AdServe{
		Url:             ad.URL,
		TrackImpression: countTrackingImpression.Count,
	}
	return &response, nil
}
