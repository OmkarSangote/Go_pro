package main

import (
	"sort"
)

func threeSum(arr []int) [][]int {
	result := [][]int{}
	sort.Ints(arr)

	for i := range arr {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		left := i + 1
		right := len(arr) - 1

		for left < right {
			sum := arr[i] + arr[left] + arr[right]

			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{arr[i], arr[left], arr[right]})
			}
		}
	}
	return result
}
