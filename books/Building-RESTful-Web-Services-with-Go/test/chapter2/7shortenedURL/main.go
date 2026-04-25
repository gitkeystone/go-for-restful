// URL 缩短服务
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jxskiss/base62"

	"github.com/gorilla/mux"
)

var db = make(map[string]string)

type URL struct {
	V string `json:"url"`
}

func NewShortenedURL(resp http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)

	var u URL
	_ = json.Unmarshal(body, &u)

	shortenedURL := base62.EncodeToString([]byte(u.V))
	longURL := u.V

	db[shortenedURL] = longURL
	_, _ = fmt.Fprintf(resp, "%s", shortenedURL)
	fmt.Printf("%v\n", db)
}

func RedirectOriginalURL(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	_, _ = fmt.Fprintf(resp, "%s", db[vars["url"]])
	fmt.Printf("%v\n", db)
	
}

func main() {
	router := mux.NewRouter()

	// POST /api/v1/new
	// {url: "/api/v1/"}
	router.HandleFunc("/api/v1/new", NewShortenedURL)

	// GET /api/v1/:url
	router.HandleFunc("/api/v1/{url}", RedirectOriginalURL)

	log.Fatal(http.ListenAndServe(":8000", router))
}
