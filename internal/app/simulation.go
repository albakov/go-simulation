package app

import (
	"bufio"
	"github.com/albakov/go-simulation/internal/app/action"
	"github.com/albakov/go-simulation/internal/app/action/action_init"
	"github.com/albakov/go-simulation/internal/app/action/action_init_grass"
	"github.com/albakov/go-simulation/internal/app/action/action_init_herbivore"
	"github.com/albakov/go-simulation/internal/app/action/action_init_predator"
	"github.com/albakov/go-simulation/internal/app/action/action_init_rock"
	"github.com/albakov/go-simulation/internal/app/action/action_init_tree"
	"github.com/albakov/go-simulation/internal/app/action/action_perform"
	"github.com/albakov/go-simulation/internal/app/action/action_perform_herbivore"
	"github.com/albakov/go-simulation/internal/app/action/action_perform_predator"
	"github.com/albakov/go-simulation/internal/app/board"
	"github.com/albakov/go-simulation/internal/app/service/menu"
	"github.com/albakov/go-simulation/internal/config"
	"github.com/albakov/go-simulation/internal/util"
	"os"
)

type Simulation struct {
	counter         int
	paused          bool
	pause           chan bool
	board           *board.Board
	initEntities    []action.Performer
	performEntities []action.Performer
	reader          *bufio.Reader
	menu            *menu.Menu
	conf            *config.Config
}

func New(conf *config.Config) *Simulation {
	s := &Simulation{
		counter:         0,
		paused:          false,
		pause:           make(chan bool),
		initEntities:    make([]action.Performer, 0),
		performEntities: make([]action.Performer, 0),
		conf:            conf,
	}

	s.board = board.New(conf.Board.Rows, conf.Board.Columns, conf.Board.DelaySeconds, &s.counter)
	s.menu = menu.New(&s.counter)

	return s
}

func (s *Simulation) StartSimulation() {
	s.menu.OnStart()
	s.addActions()
	s.initActions()

	go func() {
		s.menu.CommandListener(s.pause, &s.paused)
	}()

	go func() {
		s.menu.KeyboardListener(s.pause, &s.paused)
	}()

	for {
		s.exitIfNeed()
		s.board.DrawWorld()
		s.NextTurn()
		s.counter++
	}
}

func (s *Simulation) NextTurn() {
	for _, i := range s.performEntities {
		i.Perform()
	}
}

func (s *Simulation) addActions() {
	ai := action_init.New(s.board)
	s.initEntities = append(s.initEntities, action_init_grass.New(ai, s.conf))
	s.initEntities = append(s.initEntities, action_init_tree.New(ai, s.conf))
	s.initEntities = append(s.initEntities, action_init_rock.New(ai, s.conf))
	s.initEntities = append(s.initEntities, action_init_herbivore.New(ai, s.conf))
	s.initEntities = append(s.initEntities, action_init_predator.New(ai, s.conf))

	ap := action_perform.New(s.board, s.pause, &s.paused)
	s.performEntities = append(s.performEntities, action_perform_herbivore.New(ap))
	s.performEntities = append(s.performEntities, action_perform_predator.New(ap))
	s.performEntities = append(s.performEntities, action_init_grass.New(ai, s.conf))
}

func (s *Simulation) initActions() {
	for _, i := range s.initEntities {
		i.Perform()
	}
}

func (s *Simulation) exitIfNeed() {
	if !s.board.CanContinue() {
		util.ShowMessage(util.MessageNoReasonToContinue)
		util.ShowMessage(util.MessageCounter, s.counter)
		os.Exit(0)
	}
}
