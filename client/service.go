package client

import (
	"context"
	"fmt"

	"github.com/rrd1986/go-rabbitmq-client/logs"

	"github.com/rrd1986/go-rabbitmq-client/connection"

	"github.com/streadway/amqp"
)

// Amqp Client Type interface
type AmqpClientType interface {
	GetChannel(ctx context.Context) *amqp.Channel
	GetExchangeName(ctx context.Context) string
	GetRoutingKey(ctx context.Context) string
	DeclareRoutingTopology(ctx context.Context) error
	GetQueueName(ctx context.Context) string
	GetDeadLetterQueueName(ctx context.Context) string
	GetDeadLetterRoutingKey(ctx context.Context) string
	DeclareQueueRouting(ctx context.Context) error
	DeclareDeadLetterRouting(ctx context.Context) error
}

type AmqpClient struct {
	Channel      *amqp.Channel
	Config       *connection.Config
	ExchangeName string
	ExchangeType string
	RoutingKey   string
}

func NewAmqpClient(
	ctx context.Context,
	connection connection.AmqpConnectionType,
	exchangeName string,
	exchangeType string,
	routingKey string,
	declareRoutingTopology bool,
) (AmqpClientType, error) {
	conn, err := connection.Open(ctx)
	if err != nil {
		return &AmqpClient{}, err
	}
	channel, err := conn.Channel()

	if err != nil {
		return &AmqpClient{}, err
	}

	amqpClient := AmqpClient{
		Channel:      channel,
		Config:       connection.GetConfig(ctx),
		ExchangeName: exchangeName,
		ExchangeType: exchangeType,
		RoutingKey:   routingKey,
	}

	if declareRoutingTopology {
		err := amqpClient.DeclareRoutingTopology(ctx)

		if err != nil {
			return &amqpClient, fmt.Errorf(
				"failed to declare routing topology for exchange/routing-key %s/%s: %s",
				exchangeName,
				routingKey,
				err,
			)
		}
	}

	return &amqpClient, nil
}

func (client *AmqpClient) DeclareRoutingTopology(ctx context.Context) error {
	err0 := client.DeclareQueueRouting(ctx)
	err1 := client.DeclareDeadLetterRouting(ctx)

	for _, e := range []error{err0, err1} {
		if e != nil {
			return e
		}
	}

	return nil
}

func (client *AmqpClient) GetChannel(ctx context.Context) *amqp.Channel {
	return client.Channel
}

func (client *AmqpClient) GetExchangeName(ctx context.Context) string {
	return client.ExchangeName
}

func (client *AmqpClient) GetRoutingKey(ctx context.Context) string {
	return client.RoutingKey
}

func (client *AmqpClient) GetQueueName(ctx context.Context) string {
	return client.Config.QueueNamingStrategy.GetQueueName(client.ExchangeName, client.RoutingKey)
}

func (client *AmqpClient) GetDeadLetterQueueName(ctx context.Context) string {
	return fmt.Sprintf("%s.%s", client.GetQueueName(ctx), client.Config.DeadLetterQueueSuffix)
}

func (client *AmqpClient) GetDeadLetterRoutingKey(ctx context.Context) string {
	return client.GetQueueName(ctx)
}

func (client *AmqpClient) DeclareDeadLetterRouting(ctx context.Context) error {
	logs.Logger.Info(ctx, "Declaring amqp dead letter exchange %s", client.Config.DeadLetterExchangeName)
	err := client.Channel.ExchangeDeclare(
		client.Config.DeadLetterExchangeName,
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	deadLetterQueueName := client.GetDeadLetterQueueName(ctx)
	logs.Logger.Info(ctx, "Declaring amqp dead letter queue %s", deadLetterQueueName)
	_, err = client.Channel.QueueDeclare(
		deadLetterQueueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	deadLetterRoutingKey := client.GetDeadLetterRoutingKey(ctx)
	logs.Logger.Info(ctx, "Binding amqp dead letter queue %s to exchange %s on routing key %s",
		deadLetterQueueName,
		client.Config.DeadLetterExchangeName,
		deadLetterRoutingKey)

	return client.Channel.QueueBind(
		deadLetterQueueName,
		deadLetterRoutingKey,
		client.Config.DeadLetterExchangeName,
		false,
		nil,
	)
}

func (client *AmqpClient) DeclareQueueRouting(ctx context.Context) error {
	logs.Logger.Info(ctx, "Declaring Queue Routing for %s amqp exchange %s", client.ExchangeType, client.ExchangeName)
	err := client.Channel.ExchangeDeclare(
		client.ExchangeName,
		client.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}
	queueName := client.GetQueueName(ctx)
	logs.Logger.Info(ctx, "Declaring amqp queue %s", queueName)
	dlqOptions := amqp.Table{
		"x-dead-letter-exchange":    client.Config.DeadLetterExchangeName,
		"x-dead-letter-routing-key": queueName,
	}

	_, err = client.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		dlqOptions,
	)

	if err != nil {
		return err
	}
	logs.Logger.Info(ctx, "Binding amqp queue %s to exchange %s on routing key %s", queueName, client.ExchangeName, client.RoutingKey)
	return client.Channel.QueueBind(
		queueName,
		client.RoutingKey,
		client.ExchangeName,
		false,
		nil,
	)
}
