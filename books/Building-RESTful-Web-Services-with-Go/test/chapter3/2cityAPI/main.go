package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type city struct {
	Name string
	Area uint64
}

func postHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "POST" {
		var tempCity city
		err := json.NewDecoder(r.Body).Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(http.StatusText(http.StatusCreated)))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}
}

func main() {
	http.HandleFunc("/", postHandle)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
