package _073_time_needed_to_buy_tickets

func timeRequiredToBuyV1(tickets []int, k int) int {
	/*
		METHOD:
		Time complexity: O(n^2), где n - количество людей в очереди
		Space complexity: O(1), так как мы используем фиксированное количество доп. памяти, не зависящей от входных данных.
	*/
	time := 0 // Инициализируем время

	for tickets[k] > 0 { // Пока билеты k-го человека не кончатся
		for i := 0; i < len(tickets); i++ { // Проходим по всем билетам
			if tickets[i] > 0 { // Если у человека есть билеты
				tickets[i]-- // Купить билет
				time++       // Увеличить время
			}
			if tickets[k] == 0 { // Если билеты k-го человека кончились
				break // Выйти из цикла
			}
		}
	}

	return time // Вернуть время
}
