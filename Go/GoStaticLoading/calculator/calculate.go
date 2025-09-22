package calculator

import "fmt"

func Calculate(num ...int) {
	ans := 0
	for _, n := range num {
		ans += n
	}
	fmt.Println(ans)
}
