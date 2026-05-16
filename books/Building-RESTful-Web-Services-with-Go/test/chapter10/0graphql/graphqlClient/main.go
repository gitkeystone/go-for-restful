package main

import (
	"context"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

// Response of API
type Response struct {
	License struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"license"`
}

func main() {
	// create a client (safe to share across requests)
	client := graphql.NewClient("https://api.github.com/graphql")

	// make a request
	req := graphql.NewRequest(`
		query ($key: String!) {
			license (key: $key) {
				name
				description
			}
		}
	`)

	// set any variables
	req.Var("key", "apache-2.0")

	// authz
	token := os.Getenv("GITHUB_TOKEN")
	req.Header.Set("Authorization", "Bearer"+" "+token)

	// run it and capture the response
	var respData Response
	if err := client.Run(context.Background(), req, &respData); err != nil {
		log.Fatal(err)
	}

	// Result
	log.Println(respData.License.Description)
}
