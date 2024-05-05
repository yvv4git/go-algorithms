package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	/*
		METHOD: Hash table / Map / Dict
		Для решения этой задачи мы будем использовать словарь (map в Go) для отслеживания уникальных адресов.
		Мы будем разделять каждый адрес на локальную часть и доменную часть, а затем обработать локальную часть,
		удалив все точки и текст после символа '+'. Затем мы добавим обработанный адрес в словарь.
		В конце мы вернем количество ключей в словаре, которые будут представлять количество уникальных адресов.

		TIME COMPLEXITY: O(n), где n - количество адресов в списке, так как мы проходим по каждому адресу ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить все адреса в словаре.
	*/
	// Инициализируем словарь для отслеживания уникальных адресов
	uniqueEmails := make(map[string]bool)

	// Проходим по каждому адресу
	for _, email := range emails {
		// Разделяем адрес на локальную и доменную части
		parts := strings.Split(email, "@")
		local, domain := parts[0], parts[1]

		// Обрабатываем локальную часть
		// Удаляем все точки и текст после символа '+'
		local = strings.Split(local, "+")[0]
		local = strings.ReplaceAll(local, ".", "")

		// Формируем новый адрес и добавляем его в словарь
		newEmail := local + "@" + domain
		uniqueEmails[newEmail] = true
	}

	// Возвращаем количество ключей в словаре
	return len(uniqueEmails)
}

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails)) // Выводит: 2
}
