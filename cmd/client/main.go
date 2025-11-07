package main

import (
	"fmt"
	"log"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")

	connectionUrl := "amqp://guest:guest@localhost:5672/"
	connection, err := amqp.Dial(connectionUrl)

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	defer connection.Close()


	_, err = gamelogic.ClientWelcome()

	if err != nil {
		log.Fatalf("Failed to welcome client: %v", err)
	}
}
