package main

import "sort"

func topKfrequent(nums []int, k int) []int {
	mymap := make(map[int]int)

	for _, i := range nums {
		mymap[i]++
	}

	arr := make([]int, 0, len(mymap))

	for key := range mymap {
		arr = append(arr, key)
	}

	sort.Slice(arr, func(i, j int) bool {
		return mymap[arr[i]] > mymap[arr[j]]
	})

	return arr[:k]
}
