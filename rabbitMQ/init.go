package rabbitMQ

import (
	"fmt"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
	"go-cloud-disk/conf"
)

var RabbitMq *amqp.Connection

var RabbitMqSendEmailQueue = "send-email-queue"

func InitRabbitMq() {
	connString := strings.Join([]string{conf.RabbitMQ, "://", conf.RabbitMQUser, ":", conf.RabbitMQPassword, "@", conf.RabbitMQHost, ":", conf.RabbitMQPort, "/"}, "")
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %s", err))
	}
	RabbitMq = conn
}
