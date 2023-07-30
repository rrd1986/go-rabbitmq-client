package connection

import (
	context "context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConnectionOpenValid(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mqtConfig := NewConfig("localhost", 5673, "guest", "guest", "")
	amqpConnection, _ := NewAmqpConnection(mqtConfig)

	result, _ := amqpConnection.Open(ctx)
	assert.NotNil(t, result, "assert the error")
}

func TestConnectionOpenInvalid(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mqtConfig := NewConfig("localhost", 5672, "guest", "guest", "")
	amqpConnection, _ := NewAmqpConnection(mqtConfig)

	_, err := amqpConnection.Open(ctx)
	assert.NotNil(t, err, "assert the error")
}

func TestConnectionClose(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mqtConfig := NewConfig("localhost", 5673, "guest", "guest", "")
	amqpConnection, _ := NewAmqpConnection(mqtConfig)

	conn, _ := amqpConnection.Open(ctx)
	err := conn.Close()
	assert.Nil(t, err, "assert the error")
}

func TestConnectionGetConfig(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mqtConfig := NewConfig("localhost", 5673, "guest", "guest", "")
	amqpConnection, _ := NewAmqpConnection(mqtConfig)

	config := amqpConnection.GetConfig(ctx)
	assert.NotNil(t, config, "assert the error")
	assert.Equal(t, config.Hostname, "localhost", "assert the error")
}
