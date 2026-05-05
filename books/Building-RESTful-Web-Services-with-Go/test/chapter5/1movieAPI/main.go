package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DB struct {
	coll *mongo.Collection
}

type Movie struct {
	ID        any      `json:"id" bson:"_id,omitempty"`
	Name      string   `json:"name" bson:"name"`
	Year      int      `json:"year" bson:"year"`
	Directors []string `json:"directors" bson:"directors"`
	Writers   []string `json:"writers" bson:"writers"`
	BoxOffice `json:"boxOffice" bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `json:"budget" bson:"budget"`
	Gross  uint64 `json:"gross" bson:"gross"`
}

func (db *DB) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var movie Movie
	objectID, _ := bson.ObjectIDFromHex(vars["id"])
	err := db.coll.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(movie)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}
}

// PostMovie adds a new movie to our MongoDB collection
func (db *DB) PostMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	postBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(postBody, &movie)

	result, err := db.coll.InsertOne(context.TODO(), movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}
}

// UpdateMovie modifies the data of given resource
func (db *DB) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	objectID, _ := bson.ObjectIDFromHex(vars["id"])

	var movie Movie
	putBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(putBody, &movie)

	result, err := db.coll.UpdateOne(context.TODO(), bson.M{"_id": objectID}, bson.M{"$set": &movie})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}
}

// DeleteMovie removes the data from the db
func (db *DB) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	objectID, _ := bson.ObjectIDFromHex(vars["id"])
	result, err := db.coll.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}
}

func main() {
	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		fmt.Println("Disconnected from MongoDB")
	}()

	fmt.Println("Connected to MongoDB successfully")
	db := &DB{
		coll: client.Database("appDB").Collection("movies"),
	}

	r := mux.NewRouter()
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.GetMovie).Methods("GET")
	r.HandleFunc("/v1/movies", db.PostMovie).Methods("POST")
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.UpdateMovie).Methods("PUT")
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.DeleteMovie).Methods("DELETE")

	srv := &http.Server{
		Addr:         "127.0.0.1:8000",
		Handler:      r,
		ReadTimeout:  15e9,
		WriteTimeout: 15e9,
	}
	log.Fatal(srv.ListenAndServe())
}
