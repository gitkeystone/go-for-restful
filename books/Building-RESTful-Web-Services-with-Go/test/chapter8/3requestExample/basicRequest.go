package main

import (
	"fmt"
	"log"

	"github.com/levigross/grequests"
)

func main() {
	resp, err := grequests.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to make request:", err)
	}

	if resp.Ok {
		fmt.Println(resp.String())
	}
}
