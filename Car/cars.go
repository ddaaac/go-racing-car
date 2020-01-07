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

func GetWinnersName() []string {
	var winnersName []string
	winnerPosition := getWinnerPosition()
	for _, car := range cars {
		if car.ifAtThePosition(winnerPosition) {
			winnersName = append(winnersName, car.String())
		}
	}
	return winnersName
}

func getWinnerPosition() int {
	winnerPosition := initialPosition
	for _, car := range cars {
		winnerPosition = car.comparePosition(winnerPosition)
	}
	return winnerPosition
}
