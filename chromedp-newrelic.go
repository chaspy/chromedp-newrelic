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
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
