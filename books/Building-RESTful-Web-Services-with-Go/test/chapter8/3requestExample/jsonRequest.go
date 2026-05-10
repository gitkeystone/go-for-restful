package main

import (
	"github.com/levigross/grequests"
	"log"
)

func main() {
	resp, err := grequests.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	var returnData map[string]any
	resp.JSON(&returnData)
	log.Println(returnData)
}
