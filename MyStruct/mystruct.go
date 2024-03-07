package mystruct

import (
	"fmt"
	"math"
)

type CalcVY struct {
	X, Y float64
}

type DecimalToString float64
type IntToString int

func (c CalcVY) CalcPrint() float64 {
	return math.Sqrt(c.X*c.X + c.Y*c.Y)
	//return math.Floor(c.X*c.X + c.Y*c.Y)
}

func (c *CalcVY) Scale(f float64){
	c.X = c.X * f
	c.Y = c.Y * f
}

func (flValue DecimalToString) DecString() string {
        return fmt.Sprintf("%f", flValue)
}

func (IntValue IntToString) IntString() string {
	return fmt.Sprintf("%f", IntValue)
}

func MultipleReturn() (int, int){
	return 3,4
}
