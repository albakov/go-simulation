package action_perform_herbivore

import (
	"github.com/albakov/go-simulation/internal/app/action/action_perform"
)

type ActionHerbivore struct {
	actionPerform *action_perform.ActionPerform
}

func New(actionPerform *action_perform.ActionPerform) ActionHerbivore {
	return ActionHerbivore{actionPerform: actionPerform}
}

func (ah ActionHerbivore) Object() string {
	return "herbivore"
}

func (ah ActionHerbivore) Food() string {
	return "grass"
}

func (ah ActionHerbivore) Perform() {
	ah.actionPerform.Perform(ah)
}
