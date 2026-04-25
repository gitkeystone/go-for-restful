// TODO: 匹配和解析路径参数
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func ArticleHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req) // TODO: 解析路径参数

	// TODO: 路由器反向映射技术，生成动态 URL;
	url, _ := router.Get("articleRouter").URL("category", vars["category"], "id", vars["id"])
	fmt.Println(url.Path)

	resp.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(resp, "Category is: %v\n", vars["category"])
	_, _ = fmt.Fprintf(resp, "ID is: %v\n", vars["id"])
}

func main() {
	// TODO: 匹配路径参数;
	router.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("articleRouter")

	srv := &http.Server{
		Addr:         "127.0.0.1:8000",
		Handler:      router,
		ReadTimeout:  15e9,
		WriteTimeout: 15e9,
	}
	log.Fatal(srv.ListenAndServe())
}
