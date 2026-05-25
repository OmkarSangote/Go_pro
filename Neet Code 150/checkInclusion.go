package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var count1 [26]int
	var window [26]int

	// Build frequency for s1
	for i := 0; i < len(s1); i++ {
		count1[s1[i]-'a']++
	}

	left := 0

	for right := 0; right < len(s2); right++ {
		window[s2[right]-'a']++

		// Keep window size equal to len(s1)
		if right-left+1 > len(s1) {
			window[s2[left]-'a']--
			left++
		}

		// Compare arrays
		if window == count1 {
			return true
		}
	}

	return false
}
