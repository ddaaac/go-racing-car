package Util

import "math/rand"

func MakeRandomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
