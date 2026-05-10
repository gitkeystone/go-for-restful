package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

// Struct for holding response of repositories fetch API
type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Forks    int    `json:"forks"`
	Private  bool   `json:"private"`
}

// Structs for modelling JSON body in create Gist
type File struct {
	Content string `json:"content"`
}

type Gist struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
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

// Reads the files provided and creates Gist on github
func createGist(url string, args []string) *grequests.Response {
	// get description
	description := args[0]

	// remaining arguments are file names with path
	var fileContents = make(map[string]File)
	for i := 1; i < len(args); i++ {
		data, err := os.ReadFile(args[i])
		if err != nil {
			log.Println("Please check the filenames. Absolute path (or) same directory are allowed")
			return nil
		}
		var file File
		file.Content = string(data)
		fileContents[args[i]] = file
	}

	gist := Gist{
		Description: description,
		Public:      true,
		Files:       fileContents,
	}

	postBody, _ := json.Marshal(gist)
	ro.JSON = string(postBody)
	resp, err := grequests.Post(url, grequests.FromRequestOptions(ro))
	if err != nil {
		log.Println("Create request failed for Github API")
	}
	return resp
}

func main() {
	cmd := cli.Command{
		Version:               "1.0",
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "fetch",
				Aliases: []string{"f"},
				Usage:   "Fetch the repo details with user. [Usage]: githubAPI fetch userName",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.NArg() > 0 {
						user := cmd.Args().First()
						repoUrl := fmt.Sprintf("https://api.github.com/users/%s/repos", user)
						var repos []Repo
						resp := getStats(repoUrl, ro)
						resp.JSON(&repos)
						log.Println(repos)
					} else {
						log.Println("Please provide a user name")
					}
					return nil
				},
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Creates a gist from the given text.[Usage]: githubAPI create 'description' sample.txt ...",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.NArg() > 0 {
						postUrl := "https://api.github.com/gists"
						resp := createGist(postUrl, cmd.Args().Slice())
						log.Println(resp.String())
					} else {
						log.Println("Please give sufficient arguments. See -h to see help")
					}
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
