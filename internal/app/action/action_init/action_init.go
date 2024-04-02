package action_init

import (
	"github.com/albakov/go-simulation/internal/app/action"
	"github.com/albakov/go-simulation/internal/app/board"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
	"math"
	"math/rand"
	"time"
)

type Initializer interface {
	action.Performer
	SpawnRate() float64
	CreateEntity(c coordinate.Coordinate) entity.IEntity
}

type ActionInit struct {
	board *board.Board
	r     *rand.Rand
}

func New(board *board.Board) *ActionInit {
	return &ActionInit{
		board: board,
		r:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (ai ActionInit) Init(initializer Initializer) {
	rate := 0
	spawnRate := int(math.Round(float64(ai.board.Rows*ai.board.Columns) * initializer.SpawnRate()))

	for rate < spawnRate {
		c, hasCell := ai.getRandomCoordinate()

		if hasCell {
			ai.board.Entities[c] = initializer.CreateEntity(c)
		}

		rate++
	}
}

func (ai ActionInit) getRandomCoordinate() (coordinate.Coordinate, bool) {
	var randomCoordinate coordinate.Coordinate
	emptyCells := make([]coordinate.Coordinate, 0)

	for row := 0; row < ai.board.Rows; row++ {
		for column := 0; column < ai.board.Columns; column++ {
			c := coordinate.New(row, column)

			if _, ok := ai.board.Entities[c]; !ok {
				emptyCells = append(emptyCells, c)
			}
		}
	}

	if len(emptyCells) == 0 {
		return randomCoordinate, false
	}

	return emptyCells[ai.r.Intn(len(emptyCells))], true
}
