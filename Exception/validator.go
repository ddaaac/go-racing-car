package Exception

import (
	"../Util"
	"strings"
)

const separator = ","
const empty = ""
const intError = -1

func ValidateCarsName(carsName string) ([]string, error) {
	carsNameArray := strings.Split(carsName, separator)
	for _, carName := range carsNameArray {
		err := validateCarName(carName)
		if err != nil {
			return nil, err
		}
	}
	return carsNameArray, nil
}

func ValidateTryNumber(tryNumber string) (int, error) {
	intNumber, isPositiveInt := Util.StringToPositiveInt(tryNumber)
	if !isPositiveInt {
		return intError, notPositiveIntegerError
	}
	return intNumber, nil
}

func validateCarName(carName string) error {
	if carName == empty {
		return emptyCarNameError
	}
	for _, letter := range strings.ToLower(carName) {
		if !Util.IsAlphabet(letter) {
			return notAlphabetLetterError
		}
	}
	return nil
}
