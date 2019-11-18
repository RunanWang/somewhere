package main

import (
	"github.com/somewhere/app"
)

func main() {
	app := app.NewApp()
	app.Initialize()
	app.Run()
}
