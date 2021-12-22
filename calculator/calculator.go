package operations

type Calculator struct {
	x, y, result float64 
}

func (c *Calculator) SetX(x float64) {
	c.x = x
}
func (c *Calculator) SetY(y float64) {
	c.y = y
}
func (c Calculator) X() float64 {
	return c.x
}
func (c Calculator) Y() float64 {
	return c.y
}

func (c Calculator) Result() float64 {
	return c.result
}

func (c *Calculator) Add() {
	c.result = c.x + c.y
}

func (c *Calculator) Subtract() {
	c.result = c.x - c.y
}

func (c *Calculator) Multiply() {
	c.result = c.x * c.y
}

func (c *Calculator) Divide() {
	c.result = c.x / c.y
}