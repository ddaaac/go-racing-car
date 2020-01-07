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

func ValidateCarsName(carsName string) error {
	for _, carName := range strings.Split(carsName, separator) {
		err := validateCarName(carName)
		if err != noError {
			return err
		}
	}
	return noError
}

func ValidateTryNumber(tryNumber string) (int, error) {
	intNumber, isPositiveInt := validatePositiveInt(tryNumber)
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
		err := validateCarLetter(letter)
		if err != noError {
			return err
		}
	}
	return noError
}

func validateCarLetter(carLetter rune) error {
	if carLetter >= alphabetStart && carLetter <= alphabetEnd {
		return noError
	}
	return notAlphabetLetterError
}

func validatePositiveInt(stringNumber string) (int, bool) {
	intNumber, err := strconv.Atoi(stringNumber)
	if err != nil || intNumber <= zero {
		return intError, false
	}
	return intNumber, true
}
