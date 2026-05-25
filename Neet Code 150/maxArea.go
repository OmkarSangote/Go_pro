package main

func maxArea(input []int) int {
	left := 0
	right := len(input) - 1
	result := 0

	for left < right {
		area := (right - left) * minimum(input[left], input[right])
		result = max(area, result)

		if input[left] < input[right] {
			left++
		} else {
			right--
		}
	}
	return result

}

func minimum(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
