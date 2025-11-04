package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	connectionUrl := "amqp://guest:guest@localhost:5672/"
	connection, err := amqp.Dial(connectionUrl)

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	defer connection.Close()

	fmt.Println("Connection successful!")

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)

	done := make(chan bool, 1)

	go func() {
		sig := <- sigs
		fmt.Printf("\n Received signal: %v", sig)

		done <- true
	}()

	<-done
	fmt.Println("\n Application exited gracefully")
}
