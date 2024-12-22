//go:build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	result := randomNumber % 10
	fmt.Println("Случайное число от 0 до 9:", result)
}
