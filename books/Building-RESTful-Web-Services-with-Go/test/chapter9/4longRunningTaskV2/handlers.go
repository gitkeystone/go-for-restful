package main

import (
	"encoding/json"
	"long-running-task-v1/models"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// JobServer holds handler functions
type JobServer struct {
	Queue      amqp.Queue
	Channel    *amqp.Channel
	Connection *amqp.Connection
	RDB        *redis.Client
}

func (s *JobServer) publish(jsonBody []byte) error {
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonBody,
	}

	err := s.Channel.Publish(
		"",
		queueName,
		false,
		false,
		message,
	)
	handleError(err, "Failed to publish a message")
	return err
}

func (s *JobServer) asyncDBHandler(w http.ResponseWriter, r *http.Request) {
	jobID, err := uuid.NewRandom()
	queryParams := r.URL.Query()

	// Ex: client_time: 1569174071
	unixTime, err := strconv.ParseInt(queryParams.Get("client_time"), 10, 64)
	clientTime := time.Unix(unixTime, 0)
	handleError(err, "Failed to parse client time")

	jsonBody, err := json.Marshal(models.Job{
		ID:   jobID,
		Type: "A",
		ExtraData: models.Log{
			ClientTime: clientTime,
		},
	})
	handleError(err, "Failed to marshal json")

	if s.publish(jsonBody) == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *JobServer) asyncCallBackHandler(w http.ResponseWriter, r *http.Request) {
	jobID, err := uuid.NewRandom()
	queryParams := r.URL.Query()

	// Ex: callback: http://localhost:8080/callback
	jsonBody, err := json.Marshal(models.Job{
		ID:   jobID,
		Type: "B",
		ExtraData: models.CallBack{
			CallBack: queryParams.Get("callback"),
		},
	})
	handleError(err, "Failed to marshal json")

	if s.publish(jsonBody) == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *JobServer) asyncMailHandler(w http.ResponseWriter, r *http.Request) {
	jobID, err := uuid.NewRandom()
	queryParams := r.URL.Query()

	// Ex: email_address: <EMAIL>
	jsonBody, err := json.Marshal(models.Job{
		ID:   jobID,
		Type: "C",
		ExtraData: models.Mail{
			EmailAddress: queryParams.Get("email_address"),
		},
	})
	handleError(err, "Failed to marshal json")

	if s.publish(jsonBody) == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *JobServer) statusHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	// fetch UUID from query
	uuid := queryParams.Get("uuid")

	jobStatus := s.RDB.Get(ctx, uuid)
	status := map[string]string{
		"uuid":   uuid,
		"status": jobStatus.Val(),
	}

	response, err := json.Marshal(status)
	handleError(err, "Cannot create response for client")
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
