package main

import (
	"iam/internal/app"
	"log"
)

func main() {
	app := app.New()

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
