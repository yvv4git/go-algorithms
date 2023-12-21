package ver2

import (
	"fmt"
	"testing"
)

func TestRecentCounter_Ping(t *testing.T) {
	obj := Constructor()
	fmt.Println(obj.Ping(1))    // Вывод: 1
	fmt.Println(obj.Ping(100))  // Вывод: 2
	fmt.Println(obj.Ping(3001)) // Вывод: 3
	fmt.Println(obj.Ping(3002)) // Вывод: 3
}
