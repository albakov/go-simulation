package action_perform_predator

import (
	"github.com/albakov/go-simulation/internal/app/action/action_perform"
)

type ActionPredator struct {
	actionPerform *action_perform.ActionPerform
}

func New(actionPerform *action_perform.ActionPerform) *ActionPredator {
	return &ActionPredator{actionPerform: actionPerform}
}

func (ap ActionPredator) Object() string {
	return "predator"
}

func (ap ActionPredator) Food() string {
	return "herbivore"
}

func (ap ActionPredator) Perform() {
	ap.actionPerform.Perform(ap)
}
