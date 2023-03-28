package CucumberBDD

import "errors"

type Calculator struct {
	result int
}

func (c *Calculator) Add(x int) int {
	c.result = c.result + x
	return c.result
}

func (c *Calculator) Subtract(x int) int {
	c.result = c.result - x
	return c.result
}

func (c *Calculator) MultiplyBy(x int) int {
	c.result = c.result * x
	return c.result
}

func (c *Calculator) DivideBy(x int) (int, error) {
	if x == 0 {
		err := errors.New("zero not allowed")
		return 0, err
	}
	c.result = c.result / x
	return c.result, nil
}

func (c *Calculator) Clear() {
	c.result = 0
}

func (c *Calculator) Result() int {
	return c.result
}

func (c *Calculator) Press(x int) {
	c.result = x
}
