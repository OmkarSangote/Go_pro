package main

import (
	"fmt"
	"strconv"
)

func encode(strs []string) string {
	result := ""
	for _, str := range strs {
		result += strconv.Itoa(len(str)) + "#" + str
	}
	return result
}

func decode(s string) []string {
	result := []string{}
	i := 0

	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		lengthStr := s[i:j]
		fmt.Println("LengthStr:", lengthStr)
		length, _ := strconv.Atoi(lengthStr)
		fmt.Println("length:", length)
		start := j + 1
		end := start + length
		fmt.Println("start:", start, "end:", end)
		result = append(result, s[start:end])
		i = end
		fmt.Println("i:", i)
	}
	return result
}

func main() {
	input := []string{"hello", "world"}
	encoded := encode(input)
	fmt.Println("Encoded:", encoded)
	length := len(encoded)
	fmt.Println("Length of encoded:", length)
	decoded := decode(encoded)
	fmt.Println("decoded:", decoded)
}
