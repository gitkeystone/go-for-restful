package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"helper"
	"io"
	"log"
	"net/http"
	"utils"

	"github.com/gorilla/mux"
)

type DBClient struct {
	db *sql.DB
}

type Record struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func main() {
	db, err := helper.InitDB()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbclient := &DBClient{db: db}

	// Create a new router
	r := mux.NewRouter()

	// Attach an elegant path with handler
	r.HandleFunc("/v1/short/{encoded_string:[0-9a-zA-Z]*}", dbclient.GetOriginalURL).Methods("GET")
	r.HandleFunc("/v1/short", dbclient.GenerateShortURL).Methods("POST")

	srv := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: r,

		// Good practice: enforce timeouts for servers you create!
		ReadTimeout:  15e9,
		WriteTimeout: 15e9,
	}
	log.Fatal(srv.ListenAndServe())
}

// GenerateShortURL adds URL to DB and gives back shortened string
func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record

	postBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(postBody, &record)
	err = driver.db.QueryRow("INSERT INTO web_url(URL) VALUES ($1) RETURNING ID", record.URL).Scan(&id)
	responseMap := map[string]string{
		"encoded_string": utils.ToBase62(id),
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write([]byte(response))
	}
}

// GetOriginalURL fetches the original URL for the given encoded(short) string
func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)

	fmt.Println(vars)
	// Get ID from base62 string
	id := utils.ToBase10(vars["encoded_string"])
	err := driver.db.QueryRow("SELECT URL FROM web_url WHERE ID = $1", id).Scan(&url)
	fmt.Println(url)
	// Handle response details
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]any{
			"url": url,
		}
		response, _ := json.Marshal(responseMap)
		w.Write([]byte(response))
	}
}
