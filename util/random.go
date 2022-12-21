package util

import (
	"math/rand"
	"strconv"
)

func GenerateRandom(n int) string {
	var rand_str string
	for i := 0; i < n; i++ {
		rand_str += strconv.Itoa(rand.Intn(10))
	}
	return rand_str
}
