package Car

const initialPosition = 0

type Car struct {
	name     string
	position int
}

func newCar(name string) *Car {
	newCar := Car{name, initialPosition}
	return &newCar
}
