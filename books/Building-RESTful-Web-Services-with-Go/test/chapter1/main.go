package main

//curl -i -X GET "http://localhost:8000/fastest-mirror" # Valid request

import (
	"encoding/json"
	"fmt"
	"github.com/gitkeystone/mirror-finder/mirrors"
	"log"
	"net/http"
	"time"
)

type response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL // 任意线程占据此通道，会导致所有其他线程阻塞；从而获取最快线程称号。
				latencyChan <- latency
			}
			log.Printf("Got the best mirror: %s with latency: %s", mirrorURL, latency)
		}()
	}
	return response{<-urlChan, <-latencyChan}
}

func main() {
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter, r *http.Request) {
		response := findFastest(mirrors.MirrorList)
		respJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(respJSON)
	})
	port := ":8000"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
