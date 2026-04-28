package main

import (
	"fmt"
	"net/http"
)

// 业务逻辑处理函数
func handle(w http.ResponseWriter, r *http.Request) {
	// Business logic goes here
	fmt.Println("Executing mainHandler...")
	_, _ = w.Write([]byte("OK"))
}

// 中间件本质是一个闭包函数
// 中间件：对业务逻辑处理器的封装，返回封装过的处理器
// handler -> handler wrapper
// 中间结构：预处理 - 业务逻辑处理（处理器） - 后处理
func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//请求阶段的中间件逻辑
		fmt.Println("Executing middleware before request phase!")

		// Pass control back to the handler;
		// TODO: 转移控制权，交还给业务HTTP处理器
		originalHandler.ServeHTTP(w, r)

		//响应阶段的中间件逻辑
		fmt.Println("Executing middleware after response phase!")
	})
}

func main() {
	// HandlerFunc returns a HTTP Handler
	originalHandler := http.HandlerFunc(handle)   // 把业务逻辑处理函数，转化成业务逻辑处理器
	http.Handle("/", middleware(originalHandler)) // 封装业务处理器，扩展其他非业务逻辑
	_ = http.ListenAndServe(":8000", nil)
}
