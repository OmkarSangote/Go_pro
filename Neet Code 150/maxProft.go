package main

func maxProfit(price []int) int {
	left := 0
	right := 1
	maxProfit := 0

	for right < len(price) {
		if price[left] < price[right] {
			profit := price[right] - price[left]
			maxProfit = maximum(maxProfit, profit)

		} else {
			left = right
		}
		right++
	}
	return maxProfit
}

func maximum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
