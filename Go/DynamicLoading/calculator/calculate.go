import "fmt"

type Calculator interface {
	Calculate(num ...int)
}

type calculatorImpl struct{}

func (c *calculatorImpl) Calculate(num ...int) {
	ans := 0
	for _, n := range num {
		ans += n
	}
	fmt.Println(ans)
}

// NewCalculator function to create a new instance of Calculator
func NewCalculator() Calculator {
	return &calculatorImpl{}
}