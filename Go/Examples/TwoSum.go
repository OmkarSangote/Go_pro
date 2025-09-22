package main

import "fmt"

func twosum(num []int, target int) []int {
	l := 0
	r := len(num) - 1

	for l < r {
		curSum := num[l] + num[r]
		if curSum > target {
			r--
		} else if curSum < target {
			l++
		} else {
			return []int{l + 1, r + 1} // Assuming 1-based index output
		}
	}
	return nil // No valid pair found
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twosum(nums, target)
	fmt.Println(result) // Output: [1, 2]
}
