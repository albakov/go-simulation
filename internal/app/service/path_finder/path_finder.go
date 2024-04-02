package path_finder

import (
	"github.com/albakov/go-simulation/internal/app/board"
	"github.com/albakov/go-simulation/internal/entity"
	"github.com/albakov/go-simulation/internal/entity/coordinate"
	"math/rand"
	"slices"
	"time"
)

type PathFinder struct {
	stepsLeft int
	foodType  string
	board     *board.Board
	r         *rand.Rand
	entity    entity.IEntity
	visited   map[coordinate.Coordinate]struct{}
	q         []coordinate.Coordinate
}

func New(board *board.Board, e entity.IEntity, foodType string) *PathFinder {
	return &PathFinder{
		stepsLeft: e.GetSpeed(),
		foodType:  foodType,
		board:     board,
		entity:    e,
		r:         rand.New(rand.NewSource(time.Now().UnixNano())),
		visited:   make(map[coordinate.Coordinate]struct{}),
		q:         []coordinate.Coordinate{e.Coordinates()},
	}
}

func (pf *PathFinder) Handle() {
	for len(pf.q) > 0 && pf.stepsLeft > 0 {
		from := pf.q[0]
		pf.q = slices.Delete(pf.q, 0, 1)

		if pf.isDiedOfStarvation() {
			break
		}

		// eat nearest food
		pf.eatFood(from)

		if pf.stepsLeft <= 0 {
			break
		}

		availableMoves := pf.availableMoves(from)

		// if entity stands still, increment starvation
		if len(availableMoves) == 0 {
			pf.entity.IncrementStepsHungry()

			continue
		}

		// move to empty cell
		for _, to := range availableMoves {
			if _, ok := pf.visited[to]; !ok {
				pf.visited[to] = struct{}{}
				pf.q = append(pf.q, to)
				pf.stepsLeft--
				pf.move(from, to)
				pf.entity.MakeMove(to)

				break
			}
		}
	}
}

func (pf *PathFinder) availableMoves(c coordinate.Coordinate) []coordinate.Coordinate {
	availableMovesInsideBoard := pf.availableMovesInsideBoard(c)
	availableMoves := make([]coordinate.Coordinate, 0)

	for _, c := range availableMovesInsideBoard {
		if _, ok := pf.board.Entities[c]; !ok {
			availableMoves = append(availableMoves, c)
		}
	}

	if len(availableMoves) > 0 {
		pf.r.Shuffle(len(availableMoves), func(i, j int) {
			availableMoves[i], availableMoves[j] = availableMoves[j], availableMoves[i]
		})
	}

	return availableMoves
}

func (pf *PathFinder) eatFood(c coordinate.Coordinate) {
	for _, foodCoordinate := range pf.nearestFoodCoordinates(c) {
		if pf.stepsLeft <= 0 {
			break
		}

		if food, ok := pf.board.Entities[foodCoordinate]; ok {
			if pf.entity.Object() == "predator" {
				// trying to eat food while there are steps
				for pf.stepsLeft > 0 {
					pf.stepsLeft--
					pf.entity.Attack(food)

					if food.GetHp() <= 0 {
						pf.board.RemoveEntity(foodCoordinate)
						pf.entity.IncreaseHp(food.GiveHp())
						food = nil
						break
					}
				}
			}

			if pf.entity.Object() == "herbivore" {
				pf.stepsLeft--
				pf.board.RemoveEntity(foodCoordinate)
				pf.entity.IncreaseHp(food.GiveHp())
				food = nil
			}

			pf.board.DrawWorld()
		}
	}
}

func (pf *PathFinder) nearestFoodCoordinates(c coordinate.Coordinate) []coordinate.Coordinate {
	availableMovesInsideBoard := pf.availableMovesInsideBoard(c)
	availableMoves := make([]coordinate.Coordinate, 0)

	for _, c := range availableMovesInsideBoard {
		if item, ok := pf.board.Entities[c]; ok {
			if item.Object() == pf.foodType {
				availableMoves = append(availableMoves, c)
			}
		}
	}

	return availableMoves
}

func (pf *PathFinder) availableMovesInsideBoard(c coordinate.Coordinate) []coordinate.Coordinate {
	allMoves := []coordinate.Coordinate{
		coordinate.New(c.X-1, c.Y), // up
		coordinate.New(c.X, c.Y-1), // left
		coordinate.New(c.X, c.Y+1), // right
		coordinate.New(c.X+1, c.Y), // down

		coordinate.New(c.X-1, c.Y-1), // up-left
		coordinate.New(c.X-1, c.Y+1), // up-right
		coordinate.New(c.X+1, c.Y-1), // bottom-left
		coordinate.New(c.X+1, c.Y+1), // bottom-right
	}

	boardXMin, boardYMin := 0, 0
	boardXMax, boardYMax := pf.board.Rows, pf.board.Columns
	availableMoves := make([]coordinate.Coordinate, 0)

	for _, c := range allMoves {
		if !(c.X < boardXMin || c.X >= boardXMax || c.Y < boardYMin || c.Y >= boardYMax) {
			availableMoves = append(availableMoves, c)
		}
	}

	return availableMoves
}

func (pf *PathFinder) move(from, to coordinate.Coordinate) {
	pf.board.RemoveEntity(from)
	pf.board.AddEntity(to, pf.entity)
	pf.board.DrawWorld()
}

func (pf *PathFinder) isDiedOfStarvation() bool {
	if pf.entity.GetStepsHungry() >= pf.entity.GetMaxStepsHungryBeforeDie() {
		pf.entity.DecreaseHp(1)
	}

	if pf.entity.GetHp() <= 0 {
		pf.board.RemoveEntity(pf.entity.Coordinates())
		pf.board.DrawWorld()
		pf.entity = nil

		return true
	}

	return false
}
