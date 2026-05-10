package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cli := cli.Command{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(context.Context, *cli.Command) error {
			println("boom! I say!")
			return nil
		},
	}
	if err := cli.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
