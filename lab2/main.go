package lab2

import (
	"lab2/functions"
	"math"
)

func Function(x float64) float64 {
	if x < 0 {
		return functions.Tan(2*x) + (functions.Cos(x) / functions.Abs(functions.Ln(3*x+1)))
	} else {
		int_x := int(x)
		return functions.Pow(math.E, int_x)*functions.Sin(4*x) - functions.Sqrt(x+1)
	}
}
