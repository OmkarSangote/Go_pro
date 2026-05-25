package main

func histogramArea(num []int) int {
	n := len(num)
	area := 0

	left := make([]int, n)
	right := make([]int, n)

	stack := []int{}

	// --------- Nearest Smaller to Right ---------
	for i := n - 1; i >= 0; i-- {

		for len(stack) > 0 && num[stack[len(stack)-1]] >= num[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// Clear stack
	stack = []int{}

	// --------- Nearest Smaller to Left ---------
	for i := 0; i < n; i++ {

		for len(stack) > 0 && num[stack[len(stack)-1]] >= num[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// --------- Calculate Area ---------
	for i := 0; i < n; i++ {
		width := right[i] - left[i] - 1
		currArea := num[i] * width
		area = max(area, currArea)
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
