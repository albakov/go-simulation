package action_init_herbivore

import (
	"github.com/albakov/go-simulation/internal/app/action/action_init"
	"github.com/albakov/go-simulation/internal/config"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
	"github.com/albakov/go-simulation/internal/entity/herbivore"
)

type ActionHerbivore struct {
	spawnRate  float64
	actionInit *action_init.ActionInit
	conf       *config.Config
}

func New(actionInit *action_init.ActionInit, conf *config.Config) ActionHerbivore {
	return ActionHerbivore{
		spawnRate:  conf.Herbivore.SpawnRate,
		actionInit: actionInit,
		conf:       conf,
	}
}

func (ah ActionHerbivore) SpawnRate() float64 {
	return ah.spawnRate
}

func (ah ActionHerbivore) CreateEntity(c coordinate.Coordinate) entity.IEntity {
	return herbivore.New(c, ah.conf)
}

func (ah ActionHerbivore) Perform() {
	ah.actionInit.Init(ah)
}
