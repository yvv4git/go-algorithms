package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	// Используем функцию ReplaceAll из пакета strings для замены всех точек на "[.]"
	// ReplaceAll принимает три аргумента: исходную строку, подстроку для замены и строку, на которую заменяем
	return strings.ReplaceAll(address, ".", "[.]")
}

func main() {
	// Пример использования функции
	ipAddress := "192.168.1.1"
	defangedIP := defangIPaddr(ipAddress)
	fmt.Println(defangedIP) // Вывод: 192[.]168[.]1[.]1
}
