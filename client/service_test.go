package client

import (
	context "context"
	"fmt"
	"testing"

	"github.com/rrd1986/go-rabbitmq-client/connection"

	gomock "github.com/golang/mock/gomock"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

// amqp module does not implement any interface for the native structure bound functions it has, so its difficult to mock all subsequent calls that originate from amqp.Connection
func TestClientGetChannelWithRoutingTopology(t *testing.T) {
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

	client, _ := NewAmqpClient(
		ctx,
		mockConnection,
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		true,
	)

	result := client.GetChannel(ctx)
	assert.NotNil(t, result, "assert the error")
}

func TestClientGetChannelWithNoRoutingTopology(t *testing.T) {
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

	client, _ := NewAmqpClient(
		ctx,
		mockConnection,
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		false,
	)

	result := client.GetChannel(ctx)
	assert.NotNil(t, result, "assert the error")
}

func TestClientGetExchangeName(t *testing.T) {
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

	client, _ := NewAmqpClient(
		ctx,
		mockConnection,
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		false,
	)

	result := client.GetExchangeName(ctx)
	assert.Equal(t, result, "testing.rabbitmq", "assert the error")
}

func TestClientGetRoutingKey(t *testing.T) {
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

	client, _ := NewAmqpClient(
		ctx,
		mockConnection,
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		false,
	)

	result := client.GetRoutingKey(ctx)
	assert.Equal(t, result, "opcc.platformstatus", "assert the error")
}

func TestClientDeclareRoutingTopology(t *testing.T) {
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

	client, _ := NewAmqpClient(
		ctx,
		mockConnection,
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		false,
	)

	err := client.DeclareRoutingTopology(ctx)
	assert.Nil(t, err, "assert the error")
}

func TestClientGetQueueName(t *testing.T) {
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

	client, _ := NewAmqpClient(
		ctx,
		mockConnection,
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		false,
	)

	result := client.GetQueueName(ctx)
	assert.Equal(t, result, "testing.rabbitmq.opcc.platformstatus", "assert the error")
}

func TestClientGetDeadLetterQueueName(t *testing.T) {
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

	client, _ := NewAmqpClient(
		ctx,
		mockConnection,
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		false,
	)

	result := client.GetDeadLetterQueueName(ctx)
	assert.Equal(t, result, "testing.rabbitmq.opcc.platformstatus.dlq", "assert the error")
}

func TestClientGetDeadLetterRoutingKey(t *testing.T) {
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

	client, _ := NewAmqpClient(
		ctx,
		mockConnection,
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		false,
	)

	result := client.GetDeadLetterRoutingKey(ctx)
	assert.Equal(t, result, "testing.rabbitmq.opcc.platformstatus", "assert the error")
}

func TestClientDeclareQueueRouting(t *testing.T) {
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

	client, _ := NewAmqpClient(
		ctx,
		mockConnection,
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		false,
	)

	result := client.DeclareQueueRouting(ctx)
	assert.Nil(t, result, "assert the error")
}

func TestClientDeclareDeadLetterRouting(t *testing.T) {
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

	client, _ := NewAmqpClient(
		ctx,
		mockConnection,
		"testing.rabbitmq",
		"topic",
		"opcc.platformstatus",
		false,
	)

	result := client.DeclareDeadLetterRouting(ctx)
	assert.Nil(t, result, "assert the error")
}
