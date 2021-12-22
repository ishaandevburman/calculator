package main

import (
	"fmt"
	"gitlab.com/do12rk04/calculate_go/calculator"
)

func main() {
	calc := operations.Calculator{}
	calc.SetX(10)
	calc.SetY(1.11)
	calc.Divide()
	result := calc.Result()

	fmt.Println("Result =", result)
}