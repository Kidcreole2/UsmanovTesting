package functions

import (
	"fmt"
	"lab2"
	"lab2/functions"
	"math"
	"testing"
)

const ESP = 0.1

func TestFactorialFind(t *testing.T) {
	var tests = []struct {
		input  int
		expect float64
		out    float64
	}{
		{input: 0, expect: 1},
		{input: 1, expect: 1},
		{input: 2, expect: 2},
		{input: 5, expect: 120},
		{input: 10, expect: 3628800},
	}
	for i := 0; i < len(tests); i++ {
		tests[i].out = functions.FactorialFind(tests[i].input)
		fmt.Println("in ", tests[i].input, " out ", tests[i].out, " expect ", tests[i].expect)
		if tests[i].expect != tests[i].out {
			t.Errorf(`FactorialFind(%d) = %v, want match for %v, nil (Current test number %d, Test %d)`, tests[i].input, tests[i].out, tests[i].expect, i, len(tests))
		}
	}
}

func TestPow(t *testing.T) {
	var tests = []struct {
		input_x float64
		input_y int
		expect  float64
		out     float64
	}{
		{input_x: 2, input_y: 3, expect: 8},
		{input_x: -2, input_y: 3, expect: -8},
		{input_x: 2, input_y: -3, expect: 0.125},
		{input_x: -2, input_y: -3, expect: -0.125},
		{input_x: 0, input_y: 3, expect: 0},
		{input_x: 1, input_y: 3, expect: 1},
	}
	for i := 0; i < len(tests); i++ {
		tests[i].out = functions.Pow(tests[i].input_x, tests[i].input_y)
		fmt.Println("in x ", tests[i].input_x, " in y ", tests[i].input_y, " out ", tests[i].out, " expect ", tests[i].expect)
		if math.Abs(tests[i].expect-tests[i].out) > ESP {
			t.Errorf(`Sqr(%v,%d) = %v want match for %v, nil (Current test number %d, Test %d)`, tests[i].input_x, tests[i].input_y, tests[i].out, tests[i].expect, i, len(tests))
		}
	}
}

func TestSqrt(t *testing.T) {
	var tests = []struct {
		input  float64
		expect float64
		out    float64
	}{
		{input: 4, expect: 2},
		{input: -4, expect: -1},
		{input: 10000000000000, expect: 3162277.66017},
		{input: 1.0001, expect: 1.00005},
		{input: 0, expect: 0},
	}
	for i := 0; i < len(tests); i++ {
		tests[i].out = functions.Sqrt(tests[i].input)
		fmt.Println("in ", tests[i].input, " out ", tests[i].out, " expect ", tests[i].expect)
		if math.Abs(tests[i].expect-tests[i].out) > ESP {
			t.Errorf(`Sqrt(%v) = %v, want match for %v, nil (Current test number %d, Test %d)`, tests[i].input, tests[i].out, tests[i].expect, i, len(tests))
		}
	}
}

func TestSin(t *testing.T) {
	var tests = []struct {
		input  float64
		expect float64
		out    float64
	}{
		{input: 0, expect: 0},
		{input: math.Pi / 4, expect: 0.7071067811865475},
		{input: -math.Pi / 4, expect: -0.7071067811865475},
		{input: math.Pi / 2, expect: 1},
		{input: math.Pi, expect: 0},
		{input: 3 * math.Pi / 2, expect: -1},
	}
	for i := 0; i < len(tests); i++ {
		tests[i].out = functions.Sin(tests[i].input)
		fmt.Println("in ", tests[i].input, " out ", tests[i].out, " expect ", tests[i].expect)
		if math.Abs(tests[i].expect-tests[i].out) > ESP {
			t.Errorf(`Sin(%v) = %v, want match for %v, nil (Current test number %d, Test %d)`, tests[i].input, tests[i].out, tests[i].expect, i, len(tests))
		}
	}
}

func TestCos(t *testing.T) {
	var tests = []struct {
		input  float64
		expect float64
		out    float64
	}{
		{input: 0, expect: 1},
		{input: math.Pi / 4, expect: 0.7071067811865475},
		{input: -math.Pi / 4, expect: 0.7071067811865475},
		{input: math.Pi / 2, expect: 0},
		{input: math.Pi, expect: -1},
		{input: 3 * math.Pi / 2, expect: 0},
	}
	for i := 0; i < len(tests); i++ {
		tests[i].out = functions.Cos(tests[i].input)
		fmt.Println("in ", tests[i].input, " out ", tests[i].out, " expect ", tests[i].expect)
		if math.Abs(tests[i].expect-tests[i].out) > ESP {
			t.Errorf(`Cos(%v) = %v, want match for %v, nil (Current test number %d, Test %d)`, tests[i].input, tests[i].out, tests[i].expect, i, len(tests))
		}
	}
}

func TestTan(t *testing.T) {
	var tests = []struct {
		input  float64
		expect float64
		out    float64
	}{
		{input: 0, expect: 0},
		{input: math.Pi / 4, expect: 1},
		{input: -math.Pi / 4, expect: -1},
		{input: math.Pi / 2, expect: math.Inf(1)},
		{input: math.Pi, expect: 0},
		{input: 3 * math.Pi / 2, expect: math.Inf(-1)},
	}
	for i := 0; i < len(tests); i++ {
		tests[i].out = functions.Tan(tests[i].input)
		fmt.Println("in ", tests[i].input, " out ", tests[i].out, " expect ", tests[i].expect)
		if math.Abs(tests[i].expect-tests[i].out) > ESP {
			t.Errorf(`Tan(%v) = %v, want match for %v, nil (Current test number %d, Test %d)`, tests[i].input, tests[i].out, tests[i].expect, i, len(tests))
		}
	}
}

func TestLn(t *testing.T) {
	var tests = []struct {
		input  float64
		expect float64
		out    float64
	}{
		{input: 2, expect: 0.69314718},
		{input: -2, expect: -1},
		{input: 1.01, expect: 0.009950330594917437},
		{input: math.E, expect: 1},
		{input: 1000000, expect: 13.815510790903024},
	}
	for i := 0; i < len(tests); i++ {
		tests[i].out = functions.Ln(tests[i].input)
		fmt.Println("in ", tests[i].input, " out ", tests[i].out, " expect ", tests[i].expect)
		if math.Abs(tests[i].expect-tests[i].out) > ESP {
			t.Errorf(`Ln(%v) = %v, want match for %v, nil (Current test number %d, Test %d)`, tests[i].input, tests[i].out, tests[i].expect, i, len(tests))
		}
	}
}

func TestFunction(t *testing.T) {
	var tests = []struct {
		input  float64
		expect float64
		out    float64
	}{
		{input: -1, expect: math.Tan(-2) + (math.Cos(-1) / math.Abs(math.Log(2)))}, // x < 0
		{input: 0, expect: math.Sin(0) - math.Sqrt(1)},                             // x = 0
		{input: 1, expect: math.Exp(1)*math.Sin(4) - math.Sqrt(2)},                 // x > 0
		{input: 2, expect: math.Exp(2)*math.Sin(8) - math.Sqrt(3)},                 // x > 0
		{input: -2, expect: math.Tan(-4) + (math.Cos(-2) / math.Abs(math.Log(5)))}, // x < 0
	}
	for i := 0; i < len(tests); i++ {
		tests[i].out = lab2.Function(tests[i].input)
		fmt.Println("in ", tests[i].input, " out ", tests[i].out, " expect ", tests[i].expect)
		if math.Abs(tests[i].expect-tests[i].out) > ESP {
			t.Errorf(`Ln(%v) = %v, want match for %v, nil (Current test number %d, Test %d)`, tests[i].input, tests[i].out, tests[i].expect, i, len(tests))
		}
	}
}
