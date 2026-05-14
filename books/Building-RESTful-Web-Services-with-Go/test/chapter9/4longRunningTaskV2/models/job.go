package models

import (
	"time"

	"github.com/google/uuid"
)

// Job represents UUID of a job
type Job struct {
	ID        uuid.UUID `json:"uuid"`
	Type      string    `json:"type"`
	ExtraData any       `json:"extra_data"`
}

// Log data
type Log struct {
	ClientTime time.Time `json:"client_time"`
}

// CallBack data
type CallBack struct {
	CallBack string `json:"callback"`
}

// Mail data
type Mail struct {
	EmailAddress string `json:"email_address"`
}
