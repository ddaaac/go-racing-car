package Car

import (
	"../Util"
)

const initialPosition = 0
const minRandomValue = 0
const maxRandomValue = 9
const thresholdMoveForward = 4
const dash = "-"

var cars []*Car

func init() {
	cars = []*Car{}
}

type Car struct {
	name     string
	position int
}

func NewCars(carsName [] string) {
	for _, carName := range carsName {
		cars = append(cars, newCar(carName))
	}
}

func CarsMoveForward() []*Car {
	for _, car := range cars {
		car.moveForward()
	}
	return cars
}

func (car Car) String() string {
	var moveForwardDash string
	for i := 0; i < car.position; i++ {
		moveForwardDash += dash
	}
	return car.name + ": " + moveForwardDash
}

func newCar(name string) *Car {
	newCar := Car{name, initialPosition}
	return &newCar
}

func (car Car) moveForward() {
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
