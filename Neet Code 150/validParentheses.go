package main

func validParentheses(str string) bool {
	stack := []rune{}
	mymap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range str {
		if val, ok := mymap[ch]; ok {
			if len(stack) == 0 && stack[len(stack)-1] != val {
				return false
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}
