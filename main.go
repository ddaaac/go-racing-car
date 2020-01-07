package main

import (
	"./Car"
	"./IO"
)

func main() {
	Car.NewCars(IO.GetCarsName())
	tryNumber := IO.GetTryNumber()
	for i := 0; i < tryNumber; i++ {
		IO.PrintCarsPosition(Car.CarsMoveForward())
	}
	IO.PrintWinnersName(Car.GetWinnersName())
}
