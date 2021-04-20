package main

import (
	"connectivity-checker/app"
	"os"
)

func main() {
	app := app.Setup()

	app.Run(os.Args)
}
