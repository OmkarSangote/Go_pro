package main

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	result := []int{}

	for i := 0; i <= n-k; i++ {

		maxVal := nums[i]

		for j := i; j < i+k; j++ {
			if nums[j] > maxVal {
				maxVal = nums[j]
			}
		}

		result = append(result, maxVal)
	}

	return result
}


OR



func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	deque := []int{}      // stores indices
	result := []int{}

	for i := 0; i < len(nums); i++ {

		// Remove indices out of window
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// Remove smaller elements from back
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// Add current index
		deque = append(deque, i)

		// Window is ready
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

