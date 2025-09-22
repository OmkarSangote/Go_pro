/*
Debouncing Lab Exercise

Problem Statement:

You are building a search-as-you-type feature for a website, where the search is triggered after the user stops typing. Implement a debouncing mechanism where the search request is sent only if no new input is received for 500 milliseconds.

Requirements:
- Implement debouncing to prevent sending a search query for every keystroke.
- Simulate typing events with a delay between them.
- If there is a pause longer than 500ms between keystrokes, trigger the search.

Expected Output Example:
```
Typing: 'g'
Typing: 'go'
Typing: 'gol'
Trigger search for: 'golang'
```

---

Solution:
*/

package main

import (
	"fmt"
	"time"
)

// Simulate a search action triggered after debouncing
func search(query string) {
	fmt.Println("Trigger search for:", query)
}

func main() {
	// Simulate typing events
	input := []string{"g", "go", "gol", "gola", "golang"}
	var query string
	debounceDuration := 500 * time.Millisecond
	timer := time.NewTimer(debounceDuration)

	for _, char := range input {
		query += char
		fmt.Println("Typing:", query)
		timer.Reset(debounceDuration)

		go func(q string) {
			<-timer.C
			search(q)
		}(query)

		time.Sleep(300 * time.Millisecond) // Simulate typing delay between characters
	}
	time.Sleep(600 * time.Millisecond) // Allow last query to process after debounce period
}
