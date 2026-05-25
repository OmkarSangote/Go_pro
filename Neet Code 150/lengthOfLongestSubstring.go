package main

func lengthOfLongestSubstring(str string) int {
	left := 0
	result := 0
	mymap := make(map[byte]int)

	for right := 0; right < len(str); right++ {

		if idx, found := mymap[str[right]]; found && idx >= left {
			left = idx + 1
		}

		mymap[str[right]] = right
		result = maximum(result, right-left+1)
	}

	return result
}

func maximum(a, b int) int {
	if a > b {
		return a
	}

	return b
}
