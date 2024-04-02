package action_init_rock

import (
	"github.com/albakov/go-simulation/internal/app/action/action_init"
	"github.com/albakov/go-simulation/internal/config"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
	"github.com/albakov/go-simulation/internal/entity/rock"
)

type ActionRock struct {
	spawnRate  float64
	actionInit *action_init.ActionInit
	conf       *config.Config
}

func New(actionInit *action_init.ActionInit, conf *config.Config) ActionRock {
	return ActionRock{
		spawnRate:  conf.Rock.SpawnRate,
		actionInit: actionInit,
		conf:       conf,
	}
}

func (ar ActionRock) SpawnRate() float64 {
	return ar.spawnRate
}

func (ar ActionRock) CreateEntity(c coordinate.Coordinate) entity.IEntity {
	return rock.New(c, ar.conf)
}

func (ar ActionRock) Perform() {
	ar.actionInit.Init(ar)
}
