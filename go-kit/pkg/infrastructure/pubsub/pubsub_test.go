package pubsub

import (
	"context"
	"testing"

	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPublisher is a mock of the redisstream.Publisher
type MockPublisher struct {
	mock.Mock
}

func (m *MockPublisher) Publish(topic string, msg ...*message.Message) error {
	args := m.Called(topic, msg)
	return args.Error(0)
}

func (m *MockPublisher) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestNewRedisPub(t *testing.T) {
	// Assuming the existence of a valid redisstream.PublisherConfig
	config := redisstream.PublisherConfig{}

	// Test the creation of a new RedisPub
	redisPub := NewRedisPub(config)
	assert.NotNil(t, redisPub)
}

func TestRedisPub_Publish(t *testing.T) {
	mockPub := mocks.NewIPubRedis(t)
	mockMsg := message.NewMessage("1", []byte("test message"))

	mockPub.On("Publish", "prefix_testTopic", mock.Anything).Return(nil)

	redisPub := &RedisPub{
		globalPrefix: "prefix_",
		red:          mockPub,
	}

	err := redisPub.Publish(context.Background(), "testTopic", mockMsg)
	assert.NoError(t, err)
	mockPub.AssertExpectations(t)
}

func TestRedisPub_Close(t *testing.T) {
	mockPub := new(MockPublisher)
	mockPub.On("Close").Return(nil)

	redisPub := &RedisPub{
		red: mockPub,
	}

	err := redisPub.Close()
	assert.NoError(t, err)
	mockPub.AssertExpectations(t)
}
