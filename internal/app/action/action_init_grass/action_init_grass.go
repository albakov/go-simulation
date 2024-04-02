package action_init_grass

import (
	"github.com/albakov/go-simulation/internal/app/action/action_init"
	"github.com/albakov/go-simulation/internal/config"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
	"github.com/albakov/go-simulation/internal/entity/grass"
)

type ActionGrass struct {
	spawnRate  float64
	actionInit *action_init.ActionInit
	conf       *config.Config
}

func New(actionInit *action_init.ActionInit, conf *config.Config) ActionGrass {
	return ActionGrass{
		spawnRate:  conf.Grass.SpawnRate,
		actionInit: actionInit,
		conf:       conf,
	}
}

func (ag ActionGrass) SpawnRate() float64 {
	return ag.spawnRate
}

func (ag ActionGrass) CreateEntity(c coordinate.Coordinate) entity.IEntity {
	return grass.New(c, ag.conf)
}

func (ag ActionGrass) Perform() {
	ag.actionInit.Init(ag)
}
