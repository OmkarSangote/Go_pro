package main

func twoDSearch(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])

	l := 0
	r := (rows * cols) - 1

	for l <= r {
		mid := l + (r-l)/2

		rows = mid / 2
		cols = mid % 2

		val := matrix[rows][cols]

		if val == target {
			return true
		} else if val < target {
			l = mid + 1
		} else {
			r = r - 1
		}
	}
	return false
}
