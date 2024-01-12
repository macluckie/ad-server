package domain

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	t.Run("Create when all is ok ", func(t *testing.T) {
		ad := adProvider()
		mock := MockRepoOK{}
		service := Service{
			rp: &mock,
		}
		err := service.Create(ctx, ad)
		assert.Nil(t, err)
	})

	t.Run("Create when something wrong", func(t *testing.T) {
		ad := adProvider()
		mock := MockRepoKO{}
		service := Service{
			rp: &mock,
		}
		err := service.Create(ctx, ad)
		assert.NotNil(t, err)
		assert.Equal(t, "failed to create ad", err.Error())
	})
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	id := IdAd("test id")
	t.Run("Get when all is ok ", func(t *testing.T) {
		mock := MockRepoOK{}
		service := Service{
			rp: &mock,
		}
		ad, err := service.Get(ctx, &id)
		assert.Nil(t, err)
		assert.Equal(t, "test id", ad.Id)
	})

	t.Run("Get when something wrong", func(t *testing.T) {
		mock := MockRepoKO{}
		service := Service{
			rp: &mock,
		}
		ad, err := service.Get(ctx, &id)
		assert.NotNil(t, err)
		assert.Nil(t, ad)
	})
}

func TestServeAd(t *testing.T) {
	ctx := context.Background()
	id := IdAd("test id")
	t.Run("serveAd when all is ok ", func(t *testing.T) {
		mock := MockRepoOK{}
		service := Service{
			rp: &mock,
		}
		adServe, err := service.ServeAd(ctx, &id)
		assert.Nil(t, err)
		assert.Equal(t, int64(10), adServe.TrackImpression)
		assert.Equal(t, "http://urltest", adServe.Url)
	})

	t.Run("serveAd when something wrong", func(t *testing.T) {
		mock := MockRepoKO{}
		service := Service{
			rp: &mock,
		}
		adServe, err := service.ServeAd(ctx, &id)
		assert.NotNil(t, err)
		assert.Nil(t, adServe)
	})
}

func adProvider() *Ad {
	ad := Ad{
		Id:          "test id",
		Title:       "title test",
		Description: "description test",
		URL:         "url test",
	}
	return &ad
}

type MockRepoOK struct{}

func (m *MockRepoOK) Create(ctx context.Context, ad *Ad) error {
	return nil
}

func (m *MockRepoOK) Get(ctx context.Context, id *IdAd) (*Ad, error) {
	ad := Ad{
		Id:          "test id",
		Title:       "title test",
		Description: "description test",
		URL:         "url test",
	}
	return &ad, nil
}

func (m *MockRepoOK) ServeAd(ctx context.Context, id *IdAd) (*AdServe, error) {
	adServe := AdServe{
		Url:             "http://urltest",
		TrackImpression: 10,
	}
	return &adServe, nil
}

type MockRepoKO struct{}

func (m *MockRepoKO) Create(ctx context.Context, ad *Ad) error {
	return fmt.Errorf("failed to create ad")
}

func (m *MockRepoKO) Get(ctx context.Context, id *IdAd) (*Ad, error) {
	return nil, fmt.Errorf("failed to retrieve ad")
}

func (m *MockRepoKO) ServeAd(ctx context.Context, id *IdAd) (*AdServe, error) {
	return nil, fmt.Errorf("failed to serve ad")
}
