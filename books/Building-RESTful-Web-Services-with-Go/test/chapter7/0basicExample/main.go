package main

import (
	"helper"
	"log"
)

func main() {
	_, err := helper.InitDB()
	if err != nil {
		log.Println(err)
	}

	log.Printf("Database tables are successfully initialized.")
}
