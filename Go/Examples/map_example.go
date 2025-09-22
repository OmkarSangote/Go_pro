package main

import "fmt"

func main() {
	m := map[string]int{
		"apple":  18,
		"banana": 20,
		"cherry": 30,
	}

	fmt.Println(m)

	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	m1 := make(map[string]int)

	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m1)
	fmt.Println(m2)

	m2["foure"] = 4
	m2["five"] = 5

	fmt.Println(m2)

	appleCount := m["apple"]
	fmt.Println("Apple count: ", appleCount)

	plumCount := m["plum"]
	fmt.Println("Plum count: ", plumCount)

	value, ok := m["banana"]
	if ok {
		fmt.Println("banana count:", value)
	} else {
		fmt.Println("banana not fount")
	}

	value, ok = m["orange"]
	if ok {
		fmt.Println("orange count:", value)
	} else {
		fmt.Println("orange not fount")
	}

	delete(m, "cherry")
	fmt.Println("Map after deleting:")
	fmt.Println(m)

	//Create element from user input
}
