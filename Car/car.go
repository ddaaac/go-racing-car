package Car

import (
	"../Util"
)

const initialPosition = 0
const minRandomValue = 0
const maxRandomValue = 9
const thresholdMoveForward = 4
const dash = "-"

type Car struct {
	name     string
	position int
}

func newCar(name string) *Car {
	newCar := Car{name, initialPosition}
	return &newCar
}

func (car Car) String() string {
	var moveForwardDash string
	for i := 0; i < car.position; i++ {
		moveForwardDash += dash
	}
	return car.name + ": " + moveForwardDash
}

func (car *Car) moveForward() {
	if canMoveForward() {
		car.position++
	}
}

func canMoveForward() bool {
	randomNumber := Util.MakeRandomInt(minRandomValue, maxRandomValue)
	if randomNumber < thresholdMoveForward {
		return false
	}
	return true
}
