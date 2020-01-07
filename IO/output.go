package IO

import "fmt"
import "../Car"

func PrintLine(s string) {
	fmt.Println(s)
}

func PrintError(e error) {
	fmt.Println(e)
}

func PrintCarsPosition(cars []*Car.Car) {
	for _, car := range cars {
		printCarPosition(car)
	}
}

func printCarPosition(car *Car.Car) {
	fmt.Println(car)
}
