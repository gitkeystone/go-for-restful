package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	queueName  string = "jobQueue"
	hostString string = "127.0.0.1:8000"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func getServer(name string) JobServer {
	/*
		Create a server object and initiates
		the Channel and Queue details to publish messages.
	*/

	connection, err := amqp.Dial(os.Getenv("AMQP_URL"))
	handleError(err, "Failed to connect to RabbitMQ")

	channel, err := connection.Channel()
	handleError(err, "Failed to open a channel")

	jobQueue, err := channel.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Failed to declare a queue")

	return JobServer{
		Connection: connection,
		Channel:    channel,
		Queue:      jobQueue,
	}
}

func main() {
	jobServer := getServer(queueName)
	// Cleanup resources
	defer jobServer.Channel.Close()
	defer jobServer.Connection.Close()

	// Start Workers
	go func(connection *amqp.Connection) {
		workerProcess := Workers{
			connection: connection,
		}
		workerProcess.run()
	}(jobServer.Connection)

	// HTTP Server
	router := mux.NewRouter()
	router.HandleFunc("/job/database", jobServer.asyncDBHandler)
	router.HandleFunc("/job/callback", jobServer.asyncCallBackHandler)
	router.HandleFunc("/job/mail", jobServer.asyncMailHandler)

	httpServer := &http.Server{
		Addr:         hostString,
		Handler:      router,
		ReadTimeout:  15e9,
		WriteTimeout: 15e9,
	}
	log.Fatal(httpServer.ListenAndServe())
}
