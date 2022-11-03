// Package main provides ...
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(*cli.Context) error {
			fmt.Println("boom, I say !")
			return nil
		},
	}

	var lan string
	var port int
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "port, p",
			Value:       8888,
			Usage:       "Listenning port",
			Destination: &port,
		},

		cli.StringFlag{
			Name:        "lang, l",
			Value:       "english",
			Usage:       "read from file",
			Destination: &lan,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	// go run urfave-demo.go --port 9999 --lang go --boom
	fmt.Println("port:", port, ",lan:", lan)
}
