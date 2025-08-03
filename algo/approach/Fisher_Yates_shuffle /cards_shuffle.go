// Реальный пример использования алгоритма Фишера-Йетса - тасовка карт
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card представляет игральную карту
type Card struct {
	Suit  string
	Value string
}

func (c Card) String() string {
	return fmt.Sprintf("%s %s", c.Value, c.Suit)
}

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем колоду карт
	suits := []string{"♠", "♥", "♦", "♣"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	var deck []Card

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}

	fmt.Println("Колода до тасовки:")
	printDeck(deck[:5]) // Показываем первые 5 карт

	// Тасуем колоду с помощью алгоритма Фишера-Йетса
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}

	fmt.Println("\nКолода после тасовки:")
	printDeck(deck[:5]) // Показываем первые 5 карт после тасовки

	// Раздаем 5 карт
	fmt.Println("\nРаздача 5 карт:")
	hand := deck[:5]
	printDeck(hand)
}

// Вспомогательная функция для вывода карт
func printDeck(cards []Card) {
	for i, card := range cards {
		fmt.Printf("%d: %s\n", i+1, card)
	}
}
