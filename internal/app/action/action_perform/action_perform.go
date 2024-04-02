package action_perform

import (
	"github.com/albakov/go-simulation/internal/app/action"
	"github.com/albakov/go-simulation/internal/app/board"
	"github.com/albakov/go-simulation/internal/app/service/path_finder"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/util"
)

type Performer interface {
	action.Performer
	Object() string
	Food() string
}

type ActionPerform struct {
	board  *board.Board
	pause  chan bool
	paused *bool
}

func New(board *board.Board, pause chan bool, paused *bool) *ActionPerform {
	return &ActionPerform{board: board, pause: pause, paused: paused}
}

func (ap ActionPerform) Perform(actionPerformer Performer) {
	objects := make(map[entity.IEntity]struct{})

	for _, e := range ap.board.Entities {
		if e.Object() == actionPerformer.Object() {
			objects[e] = struct{}{}
		}
	}

	for o := range objects {
		if *ap.paused {
			util.ShowMessage(util.MessagePause)
			util.ShowMessage(util.MessagePrefix)
			<-ap.pause
		}

		path_finder.New(ap.board, o, actionPerformer.Food()).Handle()
	}
}
