package operations_test

import (
	"testing"
	"gitlab.com/do12rk04/calculate_go/calculator"
)

func TestCalculator(t *testing.T) {
	validateAdd(11,-1,10,t)
	validateSubtract(6,1,5,t)
	validateMultiply(-11,-6,66,t)
	validateDivide(62,4,15.5,t)
}

func validateAdd(x float64, y float64, expected float64 ,t *testing.T) {
	calc := operations.Calculator{}
	calc.SetX(x)
	calc.SetY(y)
	calc.Add()
	result := calc.Result()
	if result != expected {
		t.Errorf("Expected result %v, as addition of %v and %v, but got   %v", expected, x, y, result)
	}
}

func validateSubtract(x float64, y float64, expected float64 ,t *testing.T) {
	calc := operations.Calculator{}
	calc.SetX(x)
	calc.SetY(y)
	calc.Subtract()
	result := calc.Result()
	if result != expected {
		t.Errorf("Expected result %v, as subtraction of %v and %v, but got   %v", expected, x, y, result)
	}
}

func validateMultiply(x float64, y float64, expected float64 ,t *testing.T) {
	calc := operations.Calculator{}
	calc.SetX(x)
	calc.SetY(y)
	calc.Multiply()
	result := calc.Result()
	if result != expected {
		t.Errorf("Expected result %v, as multiplication of %v and %v, but got   %v", expected, x, y, result)
	}
}

func validateDivide(x float64, y float64, expected float64 ,t *testing.T) {
	calc := operations.Calculator{}
	calc.SetX(x)
	calc.SetY(y)
	calc.Divide()
	result := calc.Result()
	if result != expected {
		t.Errorf("Expected result %v, as division of %v and %v, but got   %v", expected, x, y, result)
	}
}