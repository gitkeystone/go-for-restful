package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := cli.Command{
		Version: "1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "save",
				Value: "no",
				Usage: "Should save to database (yes/no)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.NArg() > 0 {
				log.Println("Person: ", cmd.Args().First())
				log.Println("Marks: ", cmd.Args().Tail())
			}

			// Check the flag value
			if cmd.String("save") == "no" {
				log.Println("Skipping saving to the database")
			} else {
				// Add database logic here
				log.Println("Saving to the database", cmd.Args().Slice())
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
