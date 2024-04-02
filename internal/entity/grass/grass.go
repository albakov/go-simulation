package grass

import (
	"github.com/albakov/go-simulation/internal/config"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
)

type Grass struct {
	entity.Entity
	nutritionValue int
}

func New(coordinate coordinate.Coordinate, conf *config.Config) *Grass {
	return &Grass{
		Entity: entity.Entity{
			Coordinate: coordinate,
			Sign:       conf.Grass.Sign,
		},
		nutritionValue: conf.Grass.NutritionValue,
	}
}

func (g *Grass) Object() string {
	return "grass"
}

func (g *Grass) GiveHp() int {
	return g.nutritionValue
}
