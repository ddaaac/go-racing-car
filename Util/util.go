package Util

import (
	"math/rand"
	"strconv"
)

const alphabetStart = 'a'
const alphabetEnd = 'z'
const zero = 0
const intError = -1

func IsAlphabet(carLetter rune) bool {
	if carLetter < alphabetStart || carLetter > alphabetEnd {
		return false
	}
	return true
}

func StringToPositiveInt(stringNumber string) (int, bool) {
	intNumber, err := strconv.Atoi(stringNumber)
	if err != nil || intNumber <= zero {
		return intError, false
	}
	return intNumber, true
}

func MakeRandomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
