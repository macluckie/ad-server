package repo

import (
	"context"
	"fmt"
	"strconv"
	"test/grpc/config"
	"test/grpc/domain"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	client *redis.Client
}

const PREFIX_COUNTER = "tracking"

func NewRedisDb(config *config.Config) (*RedisDB, error) {
	opts, err := redis.ParseURL(config.UrlRedis)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url redis: %w", err)
	}
	db := RedisDB{
		client: redis.NewClient(opts),
	}
	ctx := context.Background()
	_, err = db.client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis, error: %w", err)
	}
	return &db, nil
}

func (r *RedisDB) GetCount(ctx context.Context, id *domain.IdAd) (*domain.CountAd, error) {
	response, err := r.client.Get(ctx, PREFIX_COUNTER+":"+string(*id)).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to query tracking Ad in database error: %w", err)
	}
	count, err := strconv.ParseInt(response, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse count, error: %w", err)
	}
	result := domain.CountAd(count)
	return &result, nil
}

func (r *RedisDB) IncrementCount(ctx context.Context, id *domain.IdAd) error {
	err := r.handleTrackingNotExist(ctx, id)
	if err != nil {
		return err
	}
	err = r.increment(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisDB) increment(ctx context.Context, id *domain.IdAd) error {
	oldTracking, err := r.client.Get(ctx, PREFIX_COUNTER+":"+string(*id)).Result()
	if err != nil {
		return err
	}
	oldValue, err := strconv.ParseInt(oldTracking, 10, 64)
	if err != nil {
		return err
	}
	err = r.client.Set(ctx, PREFIX_COUNTER+":"+string(*id), oldValue+1, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisDB) handleTrackingNotExist(ctx context.Context, id *domain.IdAd) error {
	isExist, err := r.client.Exists(ctx, PREFIX_COUNTER+":"+string(*id)).Result()
	if err != nil {
		return err
	}
	if isExist != 1 {
		err = r.client.Set(ctx, PREFIX_COUNTER+":"+string(*id), 0, 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
