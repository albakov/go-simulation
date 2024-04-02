package herbivore

import (
	"github.com/albakov/go-simulation/internal/config"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
	"github.com/albakov/go-simulation/internal/entity/creature"
)

type Herbivore struct {
	creature.Creature
	nutritionValue int
}

func New(coordinate coordinate.Coordinate, conf *config.Config) *Herbivore {
	return &Herbivore{
		Creature: creature.Creature{
			Entity: entity.Entity{
				Coordinate: coordinate,
				Sign:       conf.Herbivore.Sign,
			},
			Speed:                   conf.Herbivore.Speed,
			Hp:                      conf.Herbivore.Hp,
			StepsHungry:             0,
			MaxStepsHungryBeforeDie: conf.Herbivore.MaxStepsHungryBeforeDie,
		},
		nutritionValue: conf.Herbivore.NutritionValue,
	}
}

func (h *Herbivore) Object() string {
	return "herbivore"
}

func (h *Herbivore) GiveHp() int {
	return h.nutritionValue
}
