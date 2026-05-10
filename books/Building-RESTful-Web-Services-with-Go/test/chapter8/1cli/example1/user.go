package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "stranger",
				Usage: "your wonderful name",
			},
			&cli.IntFlag{
				Name:  "age",
				Value: 0,
				Usage: "your graceful age",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			log.Printf("Hello %s (%d years), Welcome to the command line world",
				cmd.String("name"), cmd.Int("age"))
			return nil
		},
	}

	cmd.Run(context.Background(), os.Args)
}
