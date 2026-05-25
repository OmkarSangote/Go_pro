package main

func isAscii(str byte) bool {
	return (str >= 'A' && str <= 'Z') || (str >= 'a' && str <= 'b') || (str >= '0' && str <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func validPallindrome(str string) bool {
	if len(str) == 0 {
		return false
	}

	left, right := 0, len(str)-1

	for left < right {
		for left < right && isAscii(str[left]) {
			left++
		}

		for right > left && isAscii(str[right]) {
			right--
		}

		if toLower(str[left]) != toLower(str[right]) {
			return false
		}

		left++
		right--

	}

	return true

}
