package rabbitMQ

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"go-cloud-disk/utils/logger"
)

func ConsumerMessage(ctx context.Context, queueName string) (msgs <-chan amqp.Delivery, err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		logger.Log().Error("[ConsumerMessage] Failed to open a channel: ", err)
		return nil, err
	}
	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	// mq balance
	err = ch.Qos(1, 0, false)
	if err != nil {
		logger.Log().Error("[ConsumerMessage] Failed to set Qos: ", err)
		return nil, err
	}
	return ch.Consume(q.Name, "", false, false, false, false, nil)
}
