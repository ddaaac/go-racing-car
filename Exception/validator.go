package Exception

import (
	"strings"
)

const separator = ","
const empty = ""
const alphabetStart = 'a'
const alphabetEnd = 'z'

func ValidateCarsName(carsName string) error {
	for _, carName := range strings.Split(carsName, separator) {
		err := validateCarName(carName)
		if err != noError {
			return err
		}
	}
	return noError
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
