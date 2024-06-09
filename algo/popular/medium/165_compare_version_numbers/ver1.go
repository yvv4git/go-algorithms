package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Функция для сравнения версий
func compareVersion(version1 string, version2 string) int {
	// Разделяем версии на части по точкам
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	// Проходим по каждой части версий
	for i := 0; i < len(v1) || i < len(v2); i++ {
		// Инициализируем переменные для хранения текущих частей версий
		ver1, ver2 := 0, 0

		// Если часть версии существует в первой версии, преобразуем ее в число
		if i < len(v1) {
			ver1, _ = strconv.Atoi(v1[i])
		}

		// Если часть версии существует во второй версии, преобразуем ее в число
		if i < len(v2) {
			ver2, _ = strconv.Atoi(v2[i])
		}

		// Сравниваем текущие части версий
		if ver1 > ver2 {
			// Если часть первой версии больше, возвращаем 1
			return 1
		} else if ver1 < ver2 {
			// Если часть второй версии больше, возвращаем -1
			return -1
		}
	}

	// Если все части версий оказались равны, возвращаем 0
	return 0
}

func main() {
	fmt.Println(compareVersion("1.0.1", "1.0.2")) // -1
	fmt.Println(compareVersion("1.0.2", "1.0.1")) // 1
	fmt.Println(compareVersion("1.0.1", "1.0.1")) // 0
}
