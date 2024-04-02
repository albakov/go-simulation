package action_init_predator

import (
	"github.com/albakov/go-simulation/internal/app/action/action_init"
	"github.com/albakov/go-simulation/internal/config"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
	"github.com/albakov/go-simulation/internal/entity/predator"
)

type ActionPredator struct {
	spawnRate  float64
	actionInit *action_init.ActionInit
	conf       *config.Config
}

func New(actionInit *action_init.ActionInit, conf *config.Config) ActionPredator {
	return ActionPredator{
		spawnRate:  conf.Predator.SpawnRate,
		actionInit: actionInit,
		conf:       conf,
	}
}

func (ap ActionPredator) SpawnRate() float64 {
	return ap.spawnRate
}

func (ap ActionPredator) CreateEntity(c coordinate.Coordinate) entity.IEntity {
	return predator.New(c, ap.conf)
}

func (ap ActionPredator) Perform() {
	ap.actionInit.Init(ap)
}
