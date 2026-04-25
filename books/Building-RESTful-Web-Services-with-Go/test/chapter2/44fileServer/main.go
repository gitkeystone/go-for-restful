// TODO: gorilla/mux 实现文件服务器;
// 路径前缀技术: 匹配所有以根词为前缀的 API 服务器路由
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		ReadTimeout:  15e9,
		WriteTimeout: 15e9,
	}
	log.Fatal(srv.ListenAndServe())
}
