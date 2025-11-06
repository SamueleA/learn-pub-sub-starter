package pubsub

import (
	"encoding/json"

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
