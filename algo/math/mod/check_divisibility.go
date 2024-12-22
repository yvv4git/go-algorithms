//go:build ignore

package main

import "fmt"

func main() {
	number := 15
	if number%5 == 0 {
		fmt.Println("Число делится на 5 без остатка")
	} else {
		fmt.Println("Число не делится на 5 без остатка")
	}
}
