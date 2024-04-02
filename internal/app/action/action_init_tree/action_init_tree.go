package action_init_tree

import (
	"github.com/albakov/go-simulation/internal/app/action/action_init"
	"github.com/albakov/go-simulation/internal/config"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
	"github.com/albakov/go-simulation/internal/entity/tree"
)

type ActionTree struct {
	spawnRate  float64
	actionInit *action_init.ActionInit
	conf       *config.Config
}

func New(actionInit *action_init.ActionInit, conf *config.Config) ActionTree {
	return ActionTree{
		spawnRate:  conf.Tree.SpawnRate,
		actionInit: actionInit,
		conf:       conf,
	}
}

func (at ActionTree) SpawnRate() float64 {
	return at.spawnRate
}

func (at ActionTree) CreateEntity(c coordinate.Coordinate) entity.IEntity {
	return tree.New(c, at.conf)
}

func (at ActionTree) Perform() {
	at.actionInit.Init(at)
}
