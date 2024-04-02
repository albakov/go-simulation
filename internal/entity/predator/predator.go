package predator

import (
	"github.com/albakov/go-simulation/internal/config"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
	"github.com/albakov/go-simulation/internal/entity/creature"
)

type Predator struct {
	creature.Creature
	AttackPower int
	conf        *config.Config
}

func New(coordinate coordinate.Coordinate, conf *config.Config) *Predator {
	return &Predator{
		Creature: creature.Creature{
			Entity: entity.Entity{
				Coordinate: coordinate,
				Sign:       conf.Predator.Sign,
			},
			Speed:                   conf.Predator.Speed,
			Hp:                      conf.Predator.Hp,
			StepsHungry:             0,
			MaxStepsHungryBeforeDie: conf.Predator.MaxStepsHungryBeforeDie,
		},
		AttackPower: conf.Predator.AttackPower,
	}
}

func (p *Predator) Attack(victim entity.IEntity) {
	victim.DecreaseHp(p.AttackPower)
	p.IncrementStepsHungry()
}

func (p *Predator) Object() string {
	return "predator"
}
