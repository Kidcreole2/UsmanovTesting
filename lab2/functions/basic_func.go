package functions

import (
	"fmt"
	"math"
)

func FactorialFind(num int) float64 {
	res := 1
	for i := 1; i <= num; i++ {
		res = res * i
	}
	res_float := float64(res)
	return res_float
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Pow2(x float64) float64 {
	return x * x
}

func Pow(x float64, y int) float64 {
	if y < 0 {
		x = float64(1.0 / x)
		y = -y
	}

	if y == 0 {
		return 1
	}

	if y == 1 {
		return x
	}

	if y%2 == 0 {
		return Pow(Pow2(x), y/2)
	} else {
		return x * Pow(Pow2(x), (y-1)/2)
	}
}

func Sqrt(x float64) float64 {
	if x < 0 {
		return -1
	}
	answer := x + 1
	eps := 0.00001
	for answer-x/answer > eps {
		answer = 0.5 * (answer + x/answer)
	}
	return answer
}

func Sin(x float64) float64 {
	answer := 0.0
	xx := -x * x
	pw := x
	fti := 1.0
	for i := 1; i < 25; i += 2 {
		fti /= float64(i)
		answer += pw * fti
		fti /= float64(i + 1)
		pw *= xx
	}
	fmt.Println(answer)
	return answer
}

func Cos(x float64) float64 {
	answer := 0.0
	xx := -x * x
	pw := x
	fti := 1.0
	for i := 0; i < 25; i += 2 {
		fti /= float64(i)
		answer += pw * fti
		fti /= float64(i + 1)
		pw *= xx
	}
	fmt.Println(answer)
	return answer
}

func Tan(x float64) float64 {
	return Sin(x) / Cos(x)
}

func Ln(x float64) float64 {
	var result float64 = 0

	if x <= 0 {
		return -1
	}

	if math.IsInf(x, 1) {
		return -1
	}

	if math.IsNaN(x) {
		return -1
	}

	if x > 1 {
		return Ln(x/2) - Ln(0.5)
	}

	var unit float64 = 1
	var sign float64 = 1
	var i int = 1
	for i < 1000 {
		p := Pow((x - 1), i)
		unit = p / float64(i)
		result += sign * unit

		sign = -sign
		i++
	}

	return result
}
