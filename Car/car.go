package Car

const initialPosition = 0

var cars []*Car

func init() {
	cars = []*Car{}
}

type Car struct {
	name     string
	position int
}

func NewCars(carsName[] string) {
	for _, carName := range carsName {
		cars = append(cars,
			newCar(carName))
	}
}

func newCar(name string) *Car {
	newCar := Car{name, initialPosition}
	return &newCar
}
