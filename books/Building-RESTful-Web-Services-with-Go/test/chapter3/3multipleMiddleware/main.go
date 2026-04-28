package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area uint64
}

// MIME 检查
// 中间件结构：预处理 - 处理器
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		// Filtering requests by MIME type
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, _ = w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
			return
		}

		handler.ServeHTTP(w, r) // 先处理中间件，后处理带逻辑；
	})
}

// 设置cookie
// 中间件结构：处理器 - 后处理
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r) // 先处理逻辑；后处理中间件

		// Setting cookie to every API response
		cookie := http.Cookie{
			Name:  "Server-Time(UTC)",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
		}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

func Handle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "POST" {
		var tempCity city
		err := json.NewDecoder(r.Body).Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed"))
	}
}

func main() {
	originalHandler := http.HandlerFunc(Handle)

	// 层层封装；链式调用（嵌套函数调用）；
	http.Handle("/city", filterContentType(setServerTimeCookie(originalHandler)))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
