package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandom(n int) string {
	rand.Seed(time.Now().Unix())
	length := n

	ran_str := make([]byte, length)

	// Generating Random string
	for i := 0; i < length; i++ {
		ran_str[i] = byte(48 + rand.Intn(9))
	}

	// Displaying the random string
	str := string(ran_str)
	fmt.Println(str)
	return str
}

func GenerateRandomToken(n int) string {
	rand.Seed(time.Now().Unix())
	length := n

	ran_str := make([]byte, length)

	// Generating Random string
	for i := 0; i < length; i++ {
		ran_str[i] = byte(65 + rand.Intn(25))
	}

	// Displaying the random string
	str := string(ran_str)
	fmt.Println(str)
	return str
}
