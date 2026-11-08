// Package array provides additional examples of array operations
package array

// Example: Find maximum element in an array
func FindMax(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    max := arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    return max
}

// Example: Find minimum element in an array
func FindMin(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    min := arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] < min {
            min = arr[i]
        }
    }
    return min
}

// Example: Check if array is sorted in ascending order
func IsSorted(arr []int) bool {
    for i := 1; i < len(arr); i++ {
        if arr[i] < arr[i-1] {
            return false
        }
    }
    return true
}

// Example: Reverse an array
func Reverse(arr []int) []int {
    reversed := make([]int, len(arr))
    for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
        reversed[i] = arr[j]
    }
    return reversed
}

// Example: Find the second maximum element
func FindSecondMax(arr []int) int {
    if len(arr) < 2 {
        return 0
    }
    max := arr[0]
    secondMax := 0
    for i := 1; i < len(arr); i++ {
        if arr[i] > max {
            secondMax = max
            max = arr[i]
        } else if arr[i] > secondMax && arr[i] < max {
            secondMax = arr[i]
        }
    }
    return secondMax
}

// Example: Remove duplicates from array
func RemoveDuplicates(arr []int) []int {
    if len(arr) == 0 {
        return arr
    }
    seen := make(map[int]bool)
    result := []int{}
    for _, val := range arr {
        if !seen[val] {
            seen[val] = true
            result = append(result, val)
        }
    }
    return result
}