package main

import (
	"./Car"
	"./IO"
)

func main() {
	Car.NewCars(IO.GetCarsName())
}
