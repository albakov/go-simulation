package entity

import (
	"github.com/albakov/go-simulation/internal/entity/coordinate"
	"github.com/albakov/go-simulation/internal/util"
)

type IEntity interface {
	ShowSign()
	MakeMove(to coordinate.Coordinate)
	Object() string

	Coordinates() coordinate.Coordinate
	SetCoordinates(c coordinate.Coordinate)

	GetSpeed() int

	IncreaseHp(hp int)
	DecreaseHp(hp int)
	GiveHp() int
	GetHp() int

	Attack(victim IEntity)

	IncrementStepsHungry()
	GetStepsHungry() int
	GetMaxStepsHungryBeforeDie() int
}

type Entity struct {
	coordinate.Coordinate
	Sign string
}

func New(coordinate coordinate.Coordinate) *Entity {
	return &Entity{Coordinate: coordinate, Sign: " .. "}
}

func (e *Entity) ShowSign() {
	util.ShowMessage(e.Sign)
}

func (e *Entity) MakeMove(coordinate.Coordinate) {

}

func (e *Entity) Object() string {
	return ""
}

func (e *Entity) Coordinates() coordinate.Coordinate {
	return e.Coordinate
}

func (e *Entity) SetCoordinates(c coordinate.Coordinate) {
	e.Coordinate = c
}

func (e *Entity) GetSpeed() int {
	return 0
}

func (e *Entity) IncreaseHp(int) {

}

func (e *Entity) DecreaseHp(int) {

}

func (e *Entity) GiveHp() int {
	return 0
}

func (e *Entity) GetHp() int {
	return 0
}

func (e *Entity) IncrementStepsHungry() {

}

func (e *Entity) GetStepsHungry() int {
	return 0
}

func (e *Entity) GetMaxStepsHungryBeforeDie() int {
	return 0
}

func (e *Entity) Attack(IEntity) {
}
