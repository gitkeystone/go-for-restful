package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// HealthCheck API returns date time to client
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	_, _ = io.WriteString(w, currentTime.String())
}

func main() {
	http.HandleFunc("/health", HealthCheck)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
