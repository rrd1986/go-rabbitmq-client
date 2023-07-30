//go:generate make generate

package main

import (
	"context"

	"github.com/rrd1986/go-rabbitmq-client/logs"

	"github.com/rrd1986/go-rabbitmq-client/connection"
	"github.com/rrd1986/go-rabbitmq-client/notification"
	"github.com/rrd1986/go-rabbitmq-client/producer"
	"github.com/rrd1986/go-rabbitmq-client/utils"
)

// define logger
var logger = logs.Logger

func main() {
	ctx := context.WithValue(context.Background(), "GO-RABBITMQ-CLIENT", "init")
	logger.Info(ctx, "go-rabbitmq-client Startup.")

	// Define the producer here
	utils.SetAmqpSettings()
	amqpConfig := connection.NewConfig(utils.AmqpDefaultHostname, utils.AmqpDefaultPort,
		utils.AmqpDefaultUsername, utils.AmqpDefaultPassword,
		utils.AmqpDefaultVirtualHost)

	// Create a Producer
	producerConnection, _ := connection.NewAmqpConnection(amqpConfig)
	amqpProducer, _ := producer.NewProducer(
		ctx,
		producerConnection,
		utils.JsonMessageSerializer{},
		utils.DefaultExchangeName,
		utils.ExchangeType,
		utils.RoutingKey,
		false,
	)
	// reading from the notification channel to process
	go func() {
		for msg := range notification.Messages {
			amqpProducer.PublishMessage(ctx, msg)
			logger.Info(ctx, "amqp message successfully published")
		}
	}()
}
