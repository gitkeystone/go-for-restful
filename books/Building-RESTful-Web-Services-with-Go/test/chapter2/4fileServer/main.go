// TODO: 文件服务器
package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	// Mapping to methods is possible with HttpRouter
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
