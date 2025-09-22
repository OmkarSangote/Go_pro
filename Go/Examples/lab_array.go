package main

import "fmt"

func main() {
	var sum int
	var size int
	fmt.Println("Enter the size of array: ")
	fmt.Scan(&size)
	arr := make([]int, size)
	fmt.Println("Enter the values of array: ")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	for _, value := range arr {
		if value%2 == 0 {
			sum += value
		}
	}
	fmt.Printf("Sum of even numbers: %d\n", sum)
}

/*

You are given an array (which will have a length of at least 3, but could be very large) containing integers. The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N. Write a method that takes the array as an argument and returns this "outlier" N.

Examples
[2, 4, 0, 100, 4, 11, 2602, 36] -->  11 (the only odd number)

[160, 3, 1719, 19, 11, 13, -21] --> 160 (the only even number)

func FindOutlier(integers []int) int {
  var odd, even, count int
  count = 0

    for _, value := range integers {
        if value%2 == 0 {
            count++
            even = value
        } else {
            odd = value
        }
    }

    if count == 1{
        return even
    }
    return odd
}
*/
