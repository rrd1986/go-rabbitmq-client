package connection

import (
	"github.com/rrd1986/go-rabbitmq-client/utils"
)

type Config struct {
	// Connection fields
	Hostname    string
	Port        int
	Username    string
	Password    string
	VirtualHost string

	// Routing Configuration
	DeadLetterExchangeName string
	DeadLetterQueueSuffix  string

	// Naming Configuration
	QueueNamingStrategy utils.DefaultQueueNaming
}

func NewConfig(
	hostname string,
	port int,
	username string,
	password string,
	virtualHost string,
) *Config {
	return &Config{
		Hostname:               hostname,
		Port:                   port,
		Username:               username,
		Password:               password,
		VirtualHost:            virtualHost,
		DeadLetterExchangeName: utils.DefaultDeadLetterExchangeName,
		DeadLetterQueueSuffix:  utils.DefaultDeadLetterQueueSuffix,
		QueueNamingStrategy:    utils.DefaultQueueNaming{},
	}
}
