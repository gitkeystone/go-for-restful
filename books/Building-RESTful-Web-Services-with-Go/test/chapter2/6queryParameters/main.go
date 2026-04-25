// 查询参数
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Got parameter id:%s!\n", queryParams["id"][0])
	_, _ = fmt.Fprintf(w, "Got parameter category:%s!", queryParams["category"][0])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles", QueryHandler)
	srv := &http.Server{
		Addr:         "127.0.0.1:8000",
		Handler:      r,
		ReadTimeout:  15e9,
		WriteTimeout: 15e9,
	}
	log.Fatal(srv.ListenAndServe())
}
