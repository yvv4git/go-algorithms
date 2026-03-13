package main

import "fmt"

func main() {
	var ef, mf, hf int
	var ek, mk, hk int

	// Читаем данные Феди
	fmt.Scan(&ef, &mf, &hf)
	// Читаем данные Кирилла
	fmt.Scan(&ek, &mk, &hk)

	// Время в минутах: простая = 3, средняя = 20, сложная = 120 (2 часа)
	timeFedya := ef*3 + mf*20 + hf*120
	timeKirill := ek*3 + mk*20 + hk*120

	if timeFedya > timeKirill {
		fmt.Println("Fedya")
	} else if timeKirill > timeFedya {
		fmt.Println("Kirill")
	} else {
		fmt.Println("Draw")
	}
}
