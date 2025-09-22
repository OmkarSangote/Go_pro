package main

import "fmt"

func main() {
	m := make(map[string]int)

	for {
		var fruits string
		var quantity int
		fmt.Println("Enter the fruit name or type 'quit' to exit")
		fmt.Scanln(&fruits)
		if fruits == "quit" {
			break
		}
		fmt.Println("Enter the Quantities of", fruits, ": ")
		fmt.Scanln(&quantity)

		m[fruits] = quantity
	}

	fmt.Println(m)
	var searchFruit string
	fmt.Println("Enter the fruit to be searched: ")
	fmt.Scanln(&searchFruit)
	if qty, ok := m[searchFruit]; ok {
		fmt.Println("The quantity of", searchFruit, "is", qty)
	} else {
		fmt.Println(searchFruit, "not found in the map")
	}

	var deleteFruit string
	fmt.Println("Enter the fruit name to be deleted")
	fmt.Scanln(&deleteFruit)

	if _, ok := m[deleteFruit]; ok {
		delete(m, deleteFruit)
		fmt.Println(deleteFruit, "has been deleted from the map")
		fmt.Println("The map after deletion", m)
	} else {
		fmt.Println(deleteFruit, "not found in the map")
	}

}
