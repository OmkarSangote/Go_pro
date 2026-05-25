package main

func longestConsecutive(arr []int) int {
	mymap := make(map[int]bool)
	longest := 0

	for _, i := range arr {
		mymap[i] = true
	}

	for i := range arr {

		if !mymap[i-1] {
			current := i
			length := 1

			for mymap[current+1] {
				current++
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
