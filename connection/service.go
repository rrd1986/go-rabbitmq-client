package connection

import (
	"context"
	"fmt"

	"github.com/rrd1986/go-rabbitmq-client/logs"

	"github.com/streadway/amqp"
)

// Amqp Connection Type interface
type AmqpConnectionType interface {
	GetConfig(ctx context.Context) *Config
	Open(ctx context.Context) (*amqp.Connection, error)
	Close(ctx context.Context) error
}

type AmqpConnection struct {
	Connection    *amqp.Connection
	Config        *Config
	ConnectionUrl string
}

func NewAmqpConnection(config *Config) (AmqpConnectionType, error) {
	rabbitDsn := formatamqpDsn(
		config.Hostname,
		config.Port,
		config.VirtualHost,
		config.Username,
		config.Password,
	)

	return &AmqpConnection{
		Config:        config,
		ConnectionUrl: rabbitDsn,
		Connection:    nil,
	}, nil
}

func formatamqpDsn(
	hostname string,
	port int,
	virtualHost string,
	username string,
	password string,
) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		username,
		password,
		hostname,
		port,
		virtualHost,
	)
}

func (connection *AmqpConnection) Open(ctx context.Context) (*amqp.Connection, error) {
	logs.Logger.Info(ctx, "Connecting to amqp server %s", connection.ConnectionUrl)
	conn, err := amqp.Dial(connection.ConnectionUrl)
	if err != nil {
		return nil, err
	}
	connection.Connection = conn
	return conn, nil
}

func (connection *AmqpConnection) Close(ctx context.Context) error {
	logs.Logger.Info(ctx, "Closing connection to to amqp server %s", connection.ConnectionUrl)
	err := connection.Connection.Close()
	if err != nil {
		return err
	}
	return nil
}

func (connection *AmqpConnection) GetConfig(ctx context.Context) *Config {
	logs.Logger.Info(ctx, "Configuration of amqp server %s", connection.Config)
	return connection.Config
}
