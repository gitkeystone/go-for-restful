package main

import (
	"flag"
	"log"
)

var name = flag.String("name", "stranger", "your wonderful name")

func main() {
	flag.Parse() // 命令行选项解析
	log.Printf("Hello %s, Welcome to the command line world", *name)
}
