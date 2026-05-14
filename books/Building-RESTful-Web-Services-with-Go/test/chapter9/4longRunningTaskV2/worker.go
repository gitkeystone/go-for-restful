package main

import (
	"encoding/json"
	"log"
	"long-running-task-v1/models"
	"time"

	"github.com/redis/go-redis/v9"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Workers struct {
	connection *amqp.Connection
	rdb        *redis.Client
}

func (w *Workers) run() {
	log.Printf("Workers are booted up and running")
	// Create a new redis DB
	w.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer w.connection.Close()

	channel, err := w.connection.Channel()
	handleError(err, "Failed to open a channel")
	defer channel.Close()

	jobQueue, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Failed to declare a queue")

	messages, err := channel.Consume(
		jobQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Failed to register a consumer")

	go func() {
		for message := range messages {

			job := models.Job{}
			err := json.Unmarshal(message.Body, &job)
			handleError(err, "Failed to unmarshal JSON")

			log.Printf("Workers received a message from the queue: %s", job)

			switch job.Type {
			case "A":
				w.dbWork(job)
			case "B":
				w.callbackWork(job)
			case "C":
				w.mailWork(job)
			}

			log.Printf("Received a message: %s", message.Body)
		}
	}()

	wait := make(chan bool)
	<-wait // Run long-running worker
}

func (w *Workers) dbWork(job models.Job) {
	result := job.ExtraData.(map[string]interface{})
	w.rdb.Set(ctx, job.ID.String(), "STARTED", 0)

	log.Printf("Worker %s: extracting data..., Job: %s", job.Type, result)
	w.rdb.Set(ctx, job.ID.String(), "IN PROGRESS", 0)

	time.Sleep(2e9) // 模拟数据库操作

	log.Printf("Worker %s: saving data to database..., Job: %s", job.Type, job.ID)
	w.rdb.Set(ctx, job.ID.String(), "DONE", 0)

}

func (w *Workers) callbackWork(job models.Job) {
	w.rdb.Set(ctx, job.ID.String(), "STARTED", 0)

	log.Printf("Worker %s: performing some long running process..., job: %s", job.Type, job.ID)
	w.rdb.Set(ctx, job.ID.String(), "IN PROGRESS", 0)

	time.Sleep(10e9) // 模拟回调操作

	log.Printf("Worker %s: posting the data back to the given callback..., Job: %s", job.Type, job.ID)
	w.rdb.Set(ctx, job.ID.String(), "DONE", 0)

}

func (w *Workers) mailWork(job models.Job) {
	w.rdb.Set(ctx, job.ID.String(), "STARTED", 0)

	log.Printf("Worker %s: sending the email..., Job: %s", job.Type, job.ID)
	w.rdb.Set(ctx, job.ID.String(), "IN PROGRESS", 0)

	time.Sleep(2e9)

	log.Printf("Worker %s: send the email successfully, Job: %s", job.Type, job.ID)
	w.rdb.Set(ctx, job.ID.String(), "DONE", 0)

}
