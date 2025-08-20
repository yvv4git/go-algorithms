//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	// Мапы для хранения информации о задачах
	solved := make(map[string]bool)       // Отмечаем решенные задачи
	penalty := make(map[string]int)       // Штрафное время для каждой задачи
	wrongAttempts := make(map[string]int) // Количество неверных попыток до OK

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)

		// Парсим время HH:MM
		var hours, minutes int
		fmt.Sscanf(parts[0], "%d:%d", &hours, &minutes)
		timeInMinutes := hours*60 + minutes

		// Получаем задачу и вердикт
		problem := parts[1]
		verdict := parts[2]

		// Если задача уже решена, пропускаем дальнейшую обработку
		if solved[problem] {
			continue
		}

		// Если вердикт OK
		if verdict == "OK" {
			solved[problem] = true
			// Штраф = время первой сдачи + 20 * количество неверных попыток
			penalty[problem] = timeInMinutes + 20*wrongAttempts[problem]
		} else if verdict == "WA" || verdict == "TL" || verdict == "ML" || verdict == "RE" || verdict == "SV" {
			// Увеличиваем счетчик неверных попыток, если задача еще не решена
			if !solved[problem] {
				wrongAttempts[problem]++
			}
		}
		// CE (Compilation Error) не учитывается
	}

	// Подсчитываем количество решенных задач и суммарное штрафное время
	solvedCount := len(solved)
	totalPenalty := 0
	for _, p := range penalty {
		totalPenalty += p
	}

	fmt.Printf("%d %d\n", solvedCount, totalPenalty)
}
