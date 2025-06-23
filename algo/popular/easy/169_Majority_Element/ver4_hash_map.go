package main

import "fmt"

/*
APPROACH: "Используем hash map для подсчёта количества каждого элемента. Как только какой-то элемент встретится более n/2 раз — возвращаем его."

TIME COMPLEXITY: "O(n), так как проходим по массиву один раз."

SPACE COMPLEXITY: "O(n), так как в худшем случае все элементы разные и все хранятся в map."
*/
func majorityElement(nums []int) int {
    n := len(nums)
    m := make(map[int]int)
    for _, num := range nums {
        m[num]++
        if m[num] > n/2 {
            return num
        }
    }
    return -1 // по условию задачи такого быть не может
}

func main() {
    examples := [][]int{
        {3, 2, 3},
        {2, 2, 1, 1, 1, 2, 2},
    }
    for _, nums := range examples {
        res := majorityElement(nums)
        fmt.Printf("Input: %v\n", nums)
        fmt.Printf("Majority element: %d\n\n", res)
    }
}
