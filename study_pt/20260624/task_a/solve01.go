// A. Пароли
// Подход: Для каждого пароля генерируем все его уникальные подстроки.
// Если подстрока является чьим-то паролем, добавляем количество таких
// пользователей в ответ.  В конце вычитаем единицу — самого себя, чтобы
// не учитывать пару (i,i).
// Time complexity:  O(N*L^2) — для каждого из N паролей (длины L <= 10)
//                   генерируется O(L^2) = O(100) подстрок.
// Space complexity: O(N*L^2) — хранение паролей и Map для подсчётов.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	passwords := make([]string, n)
	cnt := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &passwords[i])
		cnt[passwords[i]]++
	}

	var ans int64
	for _, s := range passwords {
		seen := make(map[string]bool)
		// Перебираем все подстроки s, пропускаем уже учтённые
		for i := 0; i < len(s); i++ {
			for j := i + 1; j <= len(s); j++ {
				sub := s[i:j]
				if seen[sub] {
					continue
				}
				seen[sub] = true
				ans += int64(cnt[sub])
			}
		}
		ans-- // исключаем пару (i, i)
	}

	fmt.Fprintln(out, ans)
}
