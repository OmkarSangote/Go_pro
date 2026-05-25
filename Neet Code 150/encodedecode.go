package main

import (
	"strconv"
)

func encode(arr []int) string {
	result := ""

	for _, i := range arr {
		result += strconv.Itoa(len(i)) + "#" + i
	}

	return result
}

func decode(s string) []string {
	result := []string{}

	i := 0

	for i < len(s) {

		j := i

		for j != '#' {
			j++
		}

		length, _ := strconv.Atoi(s[i:j])
		word := s[j+1 : j+1+length]
		result = append(result, word)

		i = j + 1 + length

	}
	return result

}

// 5#Omkar7#Sangote
