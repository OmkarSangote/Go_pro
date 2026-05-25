package main

func dailyTemperature(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	result := make([]int, n)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevIndex] = i - prevIndex

		}
		stack = append(stack, i)
	}
	return result

}
