package rock

import (
	"github.com/albakov/go-simulation/internal/config"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
)

type Rock struct {
	entity.Entity
}

func New(coordinate coordinate.Coordinate, conf *config.Config) *Rock {
	return &Rock{Entity: entity.Entity{Coordinate: coordinate, Sign: conf.Rock.Sign}}
}

func (r *Rock) Object() string {
	return "rock"
}
