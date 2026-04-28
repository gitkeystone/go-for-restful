// 分布式存储
package main

import (
	jsonparse "encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// Args holds arguments passed to JSON-RPC service
type Args struct {
	ID string
}

// Book struct holds Book JSON structure
type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

// 定义 JSON-RPC 服务
type JSONServer struct{}

// GiveBookDetail is RPC implementation
func (t *JSONServer) GiveBookDetail(_ *http.Request, args *Args, reply *Book) error {
	var books []Book

	// Read JSON file and load data
	absPath, _ := filepath.Abs("books.json")
	raw, err := os.ReadFile(absPath)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	// Unmarshal JSON raw data into books array
	err = jsonparse.Unmarshal(raw, &books)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	// Iterate over each book to find the given book
	for _, book := range books {
		// If book found, fill reply with it
		if book.ID == args.ID {
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	// Create a new RPC server
	s := rpc.NewServer()
	// Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	// Register the service by creating a new JSON server
	_ = s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	_ = http.ListenAndServe(":1234", r)
}
