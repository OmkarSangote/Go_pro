package main

func isValidSudoku(board [][]byte) bool {
	// Use array of maps to track seen digits
	rows := [9]map[byte]bool{}
	cols := [9]map[byte]bool{}
	boxes := [9]map[byte]bool{}

	for i := 0; i < 9; i++ {
		rows[i] = map[byte]bool{}
		cols[i] = map[byte]bool{}
		boxes[i] = map[byte]bool{}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}

			// Calculate box index
			boxIdx := (i/3)*3 + j/3

			if rows[i][val] {
				return false
			}
			if cols[j][val] {
				return false
			}
			if boxes[boxIdx][val] {
				return false
			}

			// Mark as seen
			rows[i][val] = true
			cols[j][val] = true
			boxes[boxIdx][val] = true
		}
	}

	return true
}
