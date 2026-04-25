// TODO: Golang 自带的Mux;
//
//	原理： 自带的 ServeMux 有个tree字段

/*
	type routingNode struct {
		// A leaf node holds a single pattern and the Handler it was registered
		// with.
		当节点是叶子节点时，这两个字段有值，表示一个完整路由的终点。
		pattern *pattern
		handler Handler

		// An interior node maps parts of the incoming request to child nodes.
		// special children keys:
		//     "/"	trailing slash (resulting from {$})
		//	   ""   single wildcard
		children   mapping[string, *routingNode]  内部节点字段
		multiChild *routingNode // child with multi wildcard 多通配符节点
		emptyChild *routingNode // optimization: child with key ""
	}
*/
package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, rand.Float64())
	})
	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, rand.Intn(100))
	})
	_ = http.ListenAndServe(":8000", newMux)
}
