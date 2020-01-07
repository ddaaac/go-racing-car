package Car

var cars []*Car

func init() {
	cars = []*Car{}
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