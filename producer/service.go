package producer

import (
	"context"
	"fmt"

	"github.com/rrd1986/go-rabbitmq-client/logs"

	"github.com/rrd1986/go-rabbitmq-client/client"
	"github.com/rrd1986/go-rabbitmq-client/connection"
	"github.com/rrd1986/go-rabbitmq-client/notification"
	"github.com/rrd1986/go-rabbitmq-client/utils"

	"github.com/streadway/amqp"
)

var Messages = make(chan notification.LzNotificationMessage)

type ProducerType interface {
	PublishMessage(ctx context.Context, msg interface{}) error
	CloseChannel(ctx context.Context) error
}

type Producer struct {
	Client            client.AmqpClientType
	MessageSerializer utils.MessageSerializer
}

func NewProducer(
	ctx context.Context,
	connection connection.AmqpConnectionType,
	serializer utils.MessageSerializer,
	exchangeName string,
	exchangeType string,
	routingKey string,
	declareRoutingTopology bool,
) (ProducerType, error) {
	amqpClient, err := client.NewAmqpClient(
		ctx,
		connection,
		exchangeName,
		exchangeType,
		routingKey,
		declareRoutingTopology,
	)

	if err != nil {
		return nil, err
	}

	logs.Logger.Info(ctx, "Initialized Producer on exchange routingkey %s...%s", exchangeName, routingKey)
	return &Producer{
		Client:            amqpClient,
		MessageSerializer: serializer,
	}, nil
}

func (producer *Producer) PublishMessage(ctx context.Context, msg interface{}) error {
	serializedMsg, err := producer.MessageSerializer.SerializeMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to serialize message for publishing: %s", err)
	}
	logs.Logger.Info(ctx, "Preparing to publish message on  %s...%s: %s",
		producer.Client.GetExchangeName(ctx),
		producer.Client.GetRoutingKey(ctx),
		string(serializedMsg))

	return producer.Client.GetChannel(ctx).Publish(
		producer.Client.GetExchangeName(ctx),
		producer.Client.GetRoutingKey(ctx),
		false,
		false,
		amqp.Publishing{ContentType: producer.MessageSerializer.GetContentType(), Body: serializedMsg},
	)
}

func (producer *Producer) CloseChannel(ctx context.Context) error {
	err := producer.Client.GetChannel(ctx).Close()
	if err != nil {
		return err
	}
	return nil
}
