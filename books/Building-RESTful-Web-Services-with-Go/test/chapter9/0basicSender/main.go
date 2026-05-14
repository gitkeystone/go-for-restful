package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

func main() {
	// Connections start with amqp.Dial() typically from a command line argument
	// or environment variable.
	connection, err := amqp.Dial(os.Getenv("AMQP_URL"))
	handleError("Failed to connect to RabbitMQ", err)

	// To cleanly shutdown by flushing kernel buffers, make sure to close and
	// wait for the response.
	defer connection.Close()

	// Most operations happen on a channel.  If any error is returned on a
	// channel, the channel will no longer be valid, throw it away and try with
	// a different channel.  If you use many channels, it's useful for the
	// server to
	channel, err := connection.Channel()
	handleError("Failed to open a channel", err)
	defer channel.Close()

	// Declare your topology here, if it doesn't exist, it will be created, if
	// it existed already and is not what you expect, then that's considered an
	// error.
	testQueue, err := channel.QueueDeclare(
		"test", // name of the queue
		true,   // message is persisted or not
		false,  // delete message when unused
		false,  // exclusive
		false,  // no waiting time
		nil,    // extra args
	)
	handleError("Failed to declare a queue", err)
	// Use your connection on this topology with either Publish or Consume, or
	// inspect your queues with QueueInspect.  It's unwise to mix Publish and
	// Consume to let TCP do its job well.
	serverTime := time.Now()
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(serverTime.String()),
	}

	err = channel.Publish(
		"",             // exchange
		testQueue.Name, // routing key(Queue)
		false,          // mandatory
		false,          // immediate
		message,
	)
	handleError("Failed to publish a message", err)
	log.Println("Successfully published a message to the queue")
}

func handleError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
