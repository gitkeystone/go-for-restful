package main

import (
	"log"
	"os"

	"github.com/levigross/grequests"
)

type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Forks    int    `json:"forks"`
	Private  bool   `json:"private"`
}

var ro = &grequests.RequestOptions{
	Auth: []string{
		os.Getenv("GITHUB_TOKEN"),
		"x-oauth-basic",
	},
}

func getStats(url string, ro *grequests.RequestOptions) *grequests.Response {
	resp, err := grequests.Get(url, grequests.FromRequestOptions(ro))
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func main() {
	var repoUrl = "https://api.github.com/users/torvalds/repos"
	resp := getStats(repoUrl, ro)

	var repos []Repo
	resp.JSON(&repos)
	log.Println(repos)
}
