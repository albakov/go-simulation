package main

import (
	"github.com/albakov/go-simulation/internal/app"
	"github.com/albakov/go-simulation/internal/config"
)

func main() {
	app.New(config.MustNew()).StartSimulation()
}
