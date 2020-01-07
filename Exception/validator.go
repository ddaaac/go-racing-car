package Exception

import (
	"strconv"
	"strings"
)

const separator = ","
const empty = ""
const alphabetStart = 'a'
const alphabetEnd = 'z'
const zero = 0
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
	intNumber, isPositiveInt := stringToPositiveInt(tryNumber)
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
		if !isAlphabet(letter) {
			return notAlphabetLetterError
		}
	}
	return nil
}

func isAlphabet(carLetter rune) bool {
	if carLetter < alphabetStart || carLetter > alphabetEnd {
		return false
	}
	return true
}

func stringToPositiveInt(stringNumber string) (int, bool) {
	intNumber, err := strconv.Atoi(stringNumber)
	if err != nil || intNumber <= zero {
		return intError, false
	}
	return intNumber, true
}
