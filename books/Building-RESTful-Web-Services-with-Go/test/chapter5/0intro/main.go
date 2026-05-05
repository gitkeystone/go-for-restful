package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Movie holds a movie data
type Movie struct {
	Name      string   `bson:"name" json:"name"`
	Year      int      `bson:"year" json:"year"`
	Directors []string `bson:"directors" json:"directors"`
	Writers   []string `bson:"writers" json:"writers"`
	BoxOffice `bson:"boxOffice" json:"boxOffice"`
}

// BoxOffice is nested in Movie
type BoxOffice struct {
	Budget uint64 `bson:"budget" json:"budget"`
	Gross  uint64 `bson:"gross" json:"gross"`
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
	coll := client.Database("appDB").Collection("movies")

	// Create a movie
	darkNight := Movie{
		Name:      "The Dark Knight",
		Year:      2008,
		Directors: []string{"Christopher Nolan"},
		Writers:   []string{"Jonathan Nolan", "Christopher Nolan"},
		BoxOffice: BoxOffice{
			Budget: 185000000,
			Gross:  533316061,
		},
	}

	_, err = coll.InsertOne(context.TODO(), darkNight)
	if err != nil {
		log.Fatal(err)
	}

	//queryResult := new(Movie)
	var queryResult bson.M
	// bson.M is used for building map for filter query
	err = coll.FindOne(context.TODO(), bson.M{"boxOffice.budget": bson.M{"$gt": 150000000}}).Decode(&queryResult)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(queryResult, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
