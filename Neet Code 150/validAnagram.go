func validAnagram (s1, s2 stirng) bool {
	if len(s1) != len(s2) {
		return false
	}

	mymap := make(map[rune]int)

	for _, i := range s1 {
		mymap[i] ++
	}

	for _, i := range s2 {
		mymap[i] --
		if mymap[i] < 0 {
			return false
		}
	} 

	return true
}


OR

func validAnagram (s1, s2 stirng) bool {
	if len(s1) != len(s2) {
		return false
	}

	a := []rune(s1)
	b := []rune(s2)

	sort.slice(a, func(i,j int)bool) {
		return a[i] < a[j]
	}

	sort.slice(b, func(i,j int)bool) {
		return b[i] < b[j]
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}