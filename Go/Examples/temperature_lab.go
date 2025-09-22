package main

import (
	"fmt"
)

var highestTemperature int = 0.0

func checkTemperature(currentTemperature int) {
	if currentTemperature > highestTemperature {
		highestTemperature = currentTemperature
		fmt.Printf("New highest temperature recorded: %d\n", highestTemperature)
	} else {
		fmt.Printf("Highest Temperature remains: %d\n", highestTemperature)
	}
}

func main() {
	var temp int
	for {
		fmt.Print("Enter the current temperature (or type 'exit' to quit): ")
		_, err := fmt.Scanf("%d", &temp)
		if err != nil {
			var exit string
			fmt.Scanf("%s", &exit)
			if exit == "exit" {
				break
			} else {
				fmt.Println("Invalid input. Please enter a valid temperature or type 'exit' to quit.")
				continue
			}
		}
		fmt.Printf("Current Temperature: %d\n", temp)
		checkTemperature(temp)
	}
}
