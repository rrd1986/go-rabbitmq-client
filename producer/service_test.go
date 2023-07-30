package producer

import (
	context "context"
	"fmt"
	"testing"

	"github.com/rrd1986/go-rabbitmq-client/connection"
	"github.com/rrd1986/go-rabbitmq-client/notification"
	"github.com/rrd1986/go-rabbitmq-client/utils"

	gomock "github.com/golang/mock/gomock"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

// amqp module does not implement any interface for the native structure bound functions it has, so its difficult to mock all subsequent calls that originate from amqp.Connection
func TestProducerPublishMessage(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", "guest", "guest", "localhost", 5673, "")
	mqtConfig := connection.NewConfig("localhost", 5673, "guest", "guest", "")
	producerConn, _ := amqp.Dial(url)
	mockConnection := connection.NewMockAmqpConnectionType(ctrl)
	mockConnection.EXPECT().
		Open(ctx).
		Return(producerConn, nil)

	mockConnection.EXPECT().
		GetConfig(ctx).
		Return(mqtConfig)

	producer, _ := NewProducer(
		ctx,
		mockConnection,
		utils.JsonMessageSerializer{},
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		true,
	)

	DummyNotificationMessage := notification.LzNotificationMessage{}
	result := producer.PublishMessage(ctx, DummyNotificationMessage)
	assert.Equal(t, result, nil, "assert the error")
}

func TestProducerCloseChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", "guest", "guest", "localhost", 5673, "")
	mqtConfig := connection.NewConfig("localhost", 5673, "guest", "guest", "")
	producerConn, _ := amqp.Dial(url)
	mockConnection := connection.NewMockAmqpConnectionType(ctrl)
	ctx := context.Background()
	mockConnection.EXPECT().
		Open(ctx).
		Return(producerConn, nil)

	mockConnection.EXPECT().
		GetConfig(ctx).
		Return(mqtConfig)

	producer, _ := NewProducer(
		ctx,
		mockConnection,
		utils.JsonMessageSerializer{},
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		true,
	)
	result := producer.CloseChannel(ctx)
	assert.Equal(t, result, nil, "assert the error")
}
