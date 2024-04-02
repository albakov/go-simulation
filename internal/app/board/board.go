package board

import (
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
	"github.com/albakov/go-simulation/internal/util"
	"os"
	"os/exec"
	"time"
)

type Board struct {
	Rows, Columns int
	delaySeconds  int
	Entities      map[coordinate.Coordinate]entity.IEntity
	counter       *int
}

func New(rows, columns, delaySeconds int, counter *int) *Board {
	return &Board{
		Rows:         rows,
		Columns:      columns,
		delaySeconds: delaySeconds,
		Entities:     make(map[coordinate.Coordinate]entity.IEntity),
		counter:      counter,
	}
}

func (b *Board) DrawWorld() {
	time.Sleep(time.Second * time.Duration(b.delaySeconds))
	b.clearScreen()

	for i := 0; i < b.Columns; i++ {
		util.ShowMessage(util.MessageBorder)
	}

	util.ShowMessage(util.MessageNewLine)

	for row := 0; row < b.Rows; row++ {
		for column := 0; column < b.Columns; column++ {
			if item, ok := b.entity(row, column); ok {
				item.ShowSign()
			} else {
				util.ShowMessage(util.MessageEmptyCell)
			}
		}

		util.ShowMessage(util.MessageNewLine)
	}

	for i := 0; i < b.Columns; i++ {
		util.ShowMessage(util.MessageBorder)
	}

	util.ShowMessage(util.MessageNewLine)
	util.ShowMessage(util.MessageCounter, *b.counter)
	util.ShowMessage(util.MessageContinueOrExit)
}

func (b *Board) CanContinue() bool {
	hasHerbivore, hasPredator := false, false

	for _, e := range b.Entities {
		if e.Object() == "herbivore" {
			hasHerbivore = true
		}

		if e.Object() == "predator" {
			hasPredator = true
		}

		if hasHerbivore && hasPredator {
			return true
		}
	}

	return false
}

func (b *Board) RemoveEntity(c coordinate.Coordinate) {
	delete(b.Entities, c)
}

func (b *Board) AddEntity(c coordinate.Coordinate, e entity.IEntity) {
	b.Entities[c] = e
}

func (b *Board) entity(x, y int) (item entity.IEntity, hasItem bool) {
	key := b.entitiesKey(x, y)
	item, hasItem = b.Entities[key]

	return
}

func (b *Board) entitiesKey(x, y int) coordinate.Coordinate {
	return coordinate.New(x, y)
}

func (b *Board) clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
