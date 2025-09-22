package main

import (
	"fmt"
)

func hasTwoOrMoreRepeated(nums []int) bool {
	counts := make(map[int]int)
	repeated := 0

	for _, num := range nums {
		counts[num]++
		if counts[num] == 2 {
			repeated++
			if repeated >= 2 {
				return true
			}
		}
	}

	return false
}

func topKFrequent(nums []int, k int) []int {
	// Step 1: Frequency map
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Bucket where index is frequency
	buckets := make([][]int, len(nums)+1) // +1 to handle freq = len(nums)
	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}

	// Step 3: Collect top K frequent from the bucket (starting from end)
	result := []int{}
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				break
			}
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(hasTwoOrMoreRepeated([]int{1, 2, 3, 2, 4, 5, 1})) // true (1 and 2 repeat)
	fmt.Println(hasTwoOrMoreRepeated([]int{1, 2, 3, 4, 5}))       // false
	fmt.Println(hasTwoOrMoreRepeated([]int{1, 1, 2, 2, 3}))       // true (1 and 2 repeat)
	fmt.Println(hasTwoOrMoreRepeated([]int{1, 1, 1, 1}))          // false (only 1 repeats)
	topKFrequent([]int{1, 2, 3, 2, 4, 5, 1}, 2)
}
