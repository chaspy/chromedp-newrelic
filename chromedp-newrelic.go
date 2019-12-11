package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "chromedp-newrelic",
		Usage: "Get metrics from NewRelic",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello chromedp-newrelic!")
			fmt.Printf("Hello %s controller\n", c.String("controller"))
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "controller",
				Aliases: []string{"c"},
				Usage:   "Specify the controller name",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
