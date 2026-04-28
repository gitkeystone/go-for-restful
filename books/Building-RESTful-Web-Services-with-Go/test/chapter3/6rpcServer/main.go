// 分布式计算
package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
}

type TimeServer int64

func (t *TimeServer) GiveServerTime(_ *Args, reply *int64) error {
	// Fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	timeServer := new(TimeServer)
	_ = rpc.Register(timeServer)
	rpc.HandleHTTP()

	// Listen for requests on port 1234
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	_ = http.Serve(l, nil)
}
