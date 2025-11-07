package pubsub

import (
	"encoding/json"
	"log"

	ctx "context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {

	jsonBytes, err := json.Marshal(val)

	if err != nil {
		return err
	}

	ch.PublishWithContext(ctx.Background(), exchange, key, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonBytes,
	})

	return nil
}

type SimpleQueueType int

const (
	durable SimpleQueueType = iota
	transient
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType,
) (*amqp.Channel, amqp.Queue, error) {
	channel, err := conn.Channel()

	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	defer channel.Close()

	queue, err := channel.QueueDeclare(
		queueName,
		queueType == durable,
		queueType == transient,
		queueType == transient,
		false,
		nil,
	)
	
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	err = channel.QueueBind(queueName, key, exchange, false, nil)

	if err != nil {
		return nil, amqp.Queue{}, err
	}

	return channel, queue, nil
}