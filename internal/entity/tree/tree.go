package tree

import (
	"github.com/albakov/go-simulation/internal/config"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
)

type Tree struct {
	entity.Entity
}

func New(coordinate coordinate.Coordinate, conf *config.Config) *Tree {
	return &Tree{Entity: entity.Entity{Coordinate: coordinate, Sign: conf.Tree.Sign}}
}

func (r *Tree) Object() string {
	return "tree"
}
