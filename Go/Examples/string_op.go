package main

import (
	"fmt"
	"strings"
)

func main() {
	name1 := "Go Programming is good and go Programming is funny"
	name9 := "Go Programming"
	fmt.Printf("Length of given string : %s is %d\n", name9, len(name9))
	fmt.Printf("Checking if %s contains : 'Go' : %t\n", name1, strings.Contains(name1, "Go"))
	name2 := strings.Replace(name1, "Programming", "Lang", 1)
	fmt.Println(name2)
	str1 := "omkar"
	str2 := "Omkar"
	str3 := "is2 Thi1s T4est 3a"
	fmt.Println(str1 == str2)

	str := "Omkar is good"
	fmt.Println(strings.HasPrefix(str, "Omkar"))
	fmt.Println(strings.HasSuffix(str, "good"))
	/
	fmt.Println(parts)
	joined := strings.Join(parts, ",")
	fmt.Println(joined)
	str4 := "  Barracuda   "
	fmt.Println(strings.TrimSpace(str4))

	big_string := "India is a large country"
	for i, value := range big_string {
		fmt.Printf("Index : %d\t  Value : %c\n", i, value)
	}

	for i, value := range big_string {
		fmt.Println(i, value)
	}

	//for reversing
	var reverse string
	for _, ch := range big_string {
		reverse = string(ch) + reverse
	}

	fmt.Println("Omkar prints reveresed : ", reverse)
	without_space := strings.ReplaceAll(reverse, " ", "")
	fmt.Println(without_space)

	str9 := "Hello, World!"
	fmt.Println(strings.ToUpper(str9))
	fmt.Println(str9[7:12])
	fmt.Println(strings.Replace(str9, "World", "Go", 1))

	str10 := "Go is fun"
	fmt.Println(str10[len(str10)-1])

	str11 := "Hello"
	fmt.Println(str11[:3])

	str12 := "banana"
	fmt.Println(strings.Count(str12, "a"))

	str13 := "Gopher"
	fmt.Println(strings.Repeat(str13, 2))

	str14 := "Hello"
	fmt.Println(strings.ReplaceAll(str14, "l", "x"))

	str15 := "omkarsangote"
	middleChar := GetMiddle(str15)
	fmt.Println("Print middle char", middleChar)

}

func GetMiddle(str string) string {
	//Code goes here!
	length := len(str)
	middle := length / 2
	if length%2 == 0 {
		return str[middle-1 : middle+1]
	} else {
		return string(str[middle])
	}

}
