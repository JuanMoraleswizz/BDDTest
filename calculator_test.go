package CucumberBDD

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatorAdd(t *testing.T) {
	c := new(Calculator)
	c.Add(3)
	assert.Equal(t, 3, c.Result())
}

func TestCalculatorSubtract(t *testing.T) {
	c := new(Calculator)
	c.Press(10)
	c.Subtract(3)
	assert.Equal(t, 7, c.Result())
}

func TestCalculatorMultiplyBy(t *testing.T) {
	c := new(Calculator)
	c.Press(10)
	c.MultiplyBy(3)
	assert.Equal(t, 30, c.Result())
}
func TestCalculatorDivideBy(t *testing.T) {
	c := new(Calculator)
	c.Press(10)
	c.DivideBy(2)

	assert.Equal(t, 5, c.Result())
}
func TestCalculatorDivideByZero(t *testing.T) {
	c := new(Calculator)
	c.Press(10)
	_, err := c.DivideBy(0)

	assert.Error(t, errors.New("zero not allowed"), err)
}
