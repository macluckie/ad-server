package domain

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCount(t *testing.T) {
	ctx := context.Background()
	id_test := IdAd("test id")
	t.Run("getCount when all is ok ", func(t *testing.T) {
		mock := MockRepoOK{}
		service := Service{
			rp: &mock,
		}
		result, err := service.GetCount(ctx, &id_test)
		assert.Nil(t, err)
		assert.Equal(t, CountAd(12), *result)
	})

	t.Run("getCount when something wrong", func(t *testing.T) {
		mock := MockRepoKO{}
		service := Service{
			rp: &mock,
		}
		result, err := service.GetCount(ctx, &id_test)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestIncrementCount(t *testing.T) {
	ctx := context.Background()
	id_test := IdAd("test id")
	t.Run("IncrementCount when all is ok ", func(t *testing.T) {
		mock := MockRepoOK{}
		service := Service{
			rp: &mock,
		}
		err := service.IncrementCount(ctx, &id_test)
		assert.Nil(t, err)

	})

	t.Run("IncrementCount when something wrong", func(t *testing.T) {
		mock := MockRepoKO{}
		service := Service{
			rp: &mock,
		}
		err := service.IncrementCount(ctx, &id_test)
		assert.NotNil(t, err)
	})

}

type MockRepoOK struct{}

func (m *MockRepoOK) GetCount(ctx context.Context, id *IdAd) (*CountAd, error) {
	test_value := CountAd(12)
	return &test_value, nil
}
func (m *MockRepoOK) IncrementCount(ctx context.Context, id *IdAd) error {
	return nil
}

type MockRepoKO struct{}

func (m *MockRepoKO) GetCount(ctx context.Context, id *IdAd) (*CountAd, error) {
	return nil, fmt.Errorf("error getCount")
}
func (m *MockRepoKO) IncrementCount(ctx context.Context, id *IdAd) error {
	return fmt.Errorf("error IncrementCount")
}
