package main

func productarrayexceptitself(arr []int) []int {

	var newarr []int
	for i := 0; i < len(arr); i++ {
		product := 1
		for j := 0; j < len(arr); j++ {
			if i != j {
				product = product * arr[j]
			}
		}
		newarr[i] = product
	}

	return newarr
}

//OR

func productarrayexceptitself(arr []int) []int {
	prefix, suffix := 1, 1
	result := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		result[i] = prefix
		prefix = prefix * arr[i]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		result[i] = result[i] * suffix
		suffix = suffix * arr[i]
	}

	return result

}
