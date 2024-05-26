package lab1

import (
	"TesingLab/lab1"
	"testing"
)

func TestSortArr(t *testing.T) {
	var tests = []struct {
		input  []int
		expect []int
		out    []int
		err    error
	}{
		{input: []int{}, expect: []int{}},
		{input: []int{1}, expect: []int{1}},
		{input: []int{1, 2, 3}, expect: []int{1, 2, 3}},
		{input: []int{3, 2, 1}, expect: []int{1, 2, 3}},
		{input: []int{1, 2, 2, 3}, expect: []int{1, 2, 2, 3}},
		{input: []int{-1, -2, -3}, expect: []int{-3, -2, -1}},
	}
	for i := 0; i < len(tests); i++ {
		tests[i].out, tests[i].err = lab1.SortArr(tests[i].input)
		for j := 0; j < len(tests[i].out); j++ {
			if tests[i].expect[j] != tests[i].out[j] {
				t.Errorf(`SortArr(%q) = %d, %v, want match for %d, nil`, tests[i].input, tests[i].out, tests[i].err, tests[i].expect)
			}
		}
	}
}

func TestPoliCheck(t *testing.T) {
	var tests = []struct {
		input  string
		expect int
		out    int
		err    error
	}{
		{input: "", expect: -1},
		{input: "kayak", expect: 0},
		{input: "racecar", expect: 0},
		{input: "hello", expect: 1},
		{input: "A man, a plan, a canal, Panama!", expect: 1},
		{input: "12321", expect: 0},
		{input: "!@#12321@#!", expect: 1},
		{input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", expect: 1},
	}
	for i := 1; i < len(tests); i++ {
		tests[i].out, tests[i].err = lab1.PoliCheck(tests[i].input)
		if tests[i].expect != tests[i].out {
			t.Errorf(`PoliCheck(%q) = %d, %v, want match for %d, nil`, tests[i].input, tests[i].out, tests[i].err, tests[i].expect)
		}
	}
}

func TestFactorialFind(t *testing.T) {
	var tests = []struct {
		input  int
		expect int
		out    int
		err    error
	}{
		{input: 0, expect: 1},
		{input: 1, expect: 1},
		{input: 2, expect: 2},
		{input: 5, expect: 120},
		{input: 10, expect: 3628800},
	}
	for i := 1; i < len(tests); i++ {
		tests[i].out, tests[i].err = lab1.FactorialFind(tests[i].input)
		if tests[i].expect != tests[i].out {
			t.Errorf(`FactorialFind(%d) = %d, %v, want match for %d, nil`, tests[i].input, tests[i].out, tests[i].err, tests[i].expect)
		}
	}
}

func TestFibonachiFind(t *testing.T) {
	var tests = []struct {
		input  int
		expect int
		out    int
		err    error
	}{
		{input: 0, expect: 0},
		{input: 1, expect: 1},
		{input: 2, expect: 1},
		{input: 5, expect: 5},
		{input: 10, expect: 55},
		{input: 92, expect: 7540113804746346429},
	}
	for i := 1; i < len(tests); i++ {
		tests[i].out, tests[i].err = lab1.FibonachiFind(tests[i].input)
		if tests[i].expect != tests[i].out {
			t.Errorf(`FibonachiFind(%d) = %d, %v, want match for %d, nil`, tests[i].input, tests[i].out, tests[i].err, tests[i].expect)
		}
	}
}

func TestDownstringFind(t *testing.T) {
	var tests = []struct {
		inputStr  string
		inputFind string
		expect    int
		out       int
		err       error
	}{
		{inputStr: "", inputFind: "", expect: -1},
		{inputStr: "abc", inputFind: "", expect: -1},
		{inputStr: "abc", inputFind: "a", expect: 0},
		{inputStr: "abc", inputFind: "bc", expect: 0},
		{inputStr: "abc", inputFind: "abd", expect: 1},
		{inputStr: "abc", inputFind: "ABC", expect: 1},
		{inputStr: "Lorem ipsum dolor sit amet", inputFind: "ipsum", expect: 0},
	}
	for i := 1; i < len(tests); i++ {
		tests[i].out, tests[i].err = lab1.DownstringFind(tests[i].inputStr, tests[i].inputFind)
		if tests[i].expect != tests[i].out {
			t.Errorf(`DownstringFind(%q,%q) = %d, %v, want match for %d, nil`, tests[i].inputStr, tests[i].inputFind, tests[i].out, tests[i].err, tests[i].expect)
		}
	}
}

func TestPrimeNumCheck(t *testing.T) {
	var tests = []struct {
		input  int
		expect int
		out    int
		err    error
	}{
		{input: 2, expect: 0},
		{input: 10, expect: 1},
		{input: 1, expect: 1},
		{input: -1, expect: -1},
		{input: 1000000007, expect: 0},
		{input: 0, expect: -1},
		{input: 100, expect: 1},
	}
	for i := 1; i < len(tests); i++ {
		tests[i].out, tests[i].err = lab1.PrimeNumCheck(tests[i].input)
		if tests[i].expect != tests[i].out {
			t.Errorf(`PrimeNumCheck(%d) = %d, %v, want match for %d, nil`, tests[i].input, tests[i].out, tests[i].err, tests[i].expect)
		}
	}
}

func TestReverceNum(t *testing.T) {
	var tests = []struct {
		input  int32
		expect int32
		out    int32
		err    error
	}{
		{input: 123, expect: 321},
		{input: -456, expect: -654},
		{input: 0, expect: 0},
		{input: 1001, expect: 1001},
		{input: 1234567890, expect: 987654321},
		{input: 2147483647, expect: 0},
		{input: 11111, expect: 11111},
	}
	for i := 1; i < len(tests); i++ {
		tests[i].out, tests[i].err = lab1.ReverceNum(tests[i].input)
		if tests[i].expect != tests[i].out {
			t.Errorf(`ReverceNum(%d) = %d, %v, want match for %d, nil`, tests[i].input, tests[i].out, tests[i].err, tests[i].expect)
		}
	}
}

func TestRomeNumConvert(t *testing.T) {
	var tests = []struct {
		input  int
		expect string
		out    string
		err    error
	}{
		{input: 0, expect: ""},
		{input: -1, expect: ""},
		{input: 4000, expect: "MMMM"},
		{input: 4, expect: "IV"},
		{input: 9, expect: "IX"},
		{input: 3999, expect: "MMMCMXCIX"},
		{input: 1984, expect: "MCMLXXXIV"},
	}
	for i := 1; i < len(tests); i++ {
		tests[i].out, tests[i].err = lab1.RomeNumConvert(tests[i].input)
		if tests[i].expect != tests[i].out {
			t.Errorf(`RomeNumConvert(%d) = %q, %v, want match for %q, nil`, tests[i].input, tests[i].out, tests[i].err, tests[i].expect)
		}
	}
}
