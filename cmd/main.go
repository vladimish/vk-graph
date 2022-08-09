package main

import (
	"github.com/vladimish/vk-graph/internal/app"
	"os"
)

func main() {
	a := app.NewApp(os.Getenv("TOKEN"))
	a.Run(1)
}
