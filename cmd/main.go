package main

import (
	"main/internal/app"
	"main/internal/config"
)

func main() {
	cfg := config.MustLoad()
	a := app.New(cfg)
	a.Run()
}
