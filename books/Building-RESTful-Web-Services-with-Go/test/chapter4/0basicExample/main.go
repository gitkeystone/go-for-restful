package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
)

func pingTime(_ *restful.Request, resp *restful.Response) {
	_, _ = io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}

func main() {
	// Create a web service
	ws := new(restful.WebService)

	// Create a route and attach it to handler in the service
	ws.Route(ws.GET("/ping").To(pingTime))

	// Add the service to application
	restful.Add(ws)

	log.Fatal(http.ListenAndServe(":8000", nil)) // 因为 restful 默认使用 http.DefaultServeMux
}
