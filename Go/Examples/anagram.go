package main

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make(map[rune]int)

	for _, ch := range s1 {
		counts[ch]++
	}

	for _, ch := range s2 {
		counts[ch]--
		if counts[ch] < 0 {
			return false
		}
	}

	return true
}
