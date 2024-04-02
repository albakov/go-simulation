package creature

import (
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
)

type Creature struct {
	entity.Entity
	Speed, Hp, StepsHungry, MaxStepsHungryBeforeDie int
}

func (c *Creature) MakeMove(to coordinate.Coordinate) {
	c.SetCoordinates(to)
	c.IncrementStepsHungry()
}

func (c *Creature) IncreaseHp(hp int) {
	c.Hp += hp
	c.StepsHungry = 0
}

func (c *Creature) DecreaseHp(hp int) {
	c.Hp -= hp
}

func (c *Creature) GetHp() int {
	return c.Hp
}

func (c *Creature) GetSpeed() int {
	return c.Speed
}

func (c *Creature) IncrementStepsHungry() {
	c.StepsHungry += 1
}

func (c *Creature) GetStepsHungry() int {
	return c.StepsHungry
}

func (c *Creature) GetMaxStepsHungryBeforeDie() int {
	return c.MaxStepsHungryBeforeDie
}
