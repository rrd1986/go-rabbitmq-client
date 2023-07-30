package utils

import "fmt"

type QueueNaming interface {
	GetQueueName(exchangeName string, routingKey string) string
}

type DefaultQueueNaming struct{}

func (DefaultQueueNaming) GetQueueName(exchangeName string, routingKey string) string {
	return fmt.Sprintf("%s.%s", exchangeName, routingKey)
}
