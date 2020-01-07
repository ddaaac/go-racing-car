package IO

import (
	"fmt"
	"strings"
)
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

func PrintWinnersName(winnersName []string) {
	fmt.Print(strings.Join(winnersName, ", "))
	fmt.Println("이/가 최종 우승했습니다")
}

func printCarPosition(car *Car.Car) {
	fmt.Print(car)
	fmt.Println(": " + car.GetDashLine())
}