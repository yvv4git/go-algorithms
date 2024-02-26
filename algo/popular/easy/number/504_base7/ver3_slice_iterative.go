package main

func convertToBase7V4(num int) string {
	if num == 0 {
		return "0"
	}

	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num // Делаем число положительным
	}

	result := make([]byte, 0)
	for num > 0 {
		result = append([]byte{byte('0' + num%7)}, result...) // Добавляем последнюю цифру в начало строки
		num /= 7                                              // Удаляем последнюю цифру из числа
	}

	if isNegative {
		result = append([]byte{'-'}, result...) // Добавляем знак "-" в начало, если число было отрицательным
	}

	return string(result)
}

//func main() {
//	num := 100
//	fmt.Println(convertToBase7(num)) // Вывод: "202"
//}
