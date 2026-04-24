package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

// UUID is a custom multiplexer
type UUID struct{}

func (u UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandomUUID(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandomUUID(w http.ResponseWriter, _ *http.Request) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b) // 填充字符
	if err != nil {
		panic(err)
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("%x", b))
}

func main() {
	mux := &UUID{}
	_ = http.ListenAndServe(":8000", mux)
}
