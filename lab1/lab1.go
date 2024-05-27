package lab1

import (
	"errors"
)

type rome_number struct {
	name  string
	value int
}

func SortArr(array []int) ([]int, error) {
	if array == nil {
		return nil, errors.New("empty value")
	}
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				tmp := array[j]
				array[j] = array[i]
				array[i] = tmp
			}
		}
	}
	return array, nil
}

func PoliCheck(str string) (int, error) {
	if str == "" {
		return -1, errors.New("empty value")
	}
	res := 0
	i := 0
	j := len(str) - 1
	for i < j {
		if str[i] != str[j] {
			res = 1
		}
		i++
		j--
	}
	return res, nil
}

func FactorialFind(num int) (int, error) {
	if num < 1 {
		return -1, errors.New("value must be > 0")
	}
	res := 1
	for i := 1; i <= num; i++ {
		res = res * i
	}
	return res, nil
}

func FibonachiFind(num int) (int, error) {
	if num < 0 {
		return -1, errors.New("value must be >= 0")
	}
	res := 1
	prev_res := 0
	for i := 1; i < num; i++ {
		tmp := res
		res += prev_res
		prev_res = tmp
	}
	return res, nil
}
func DownstringFind(str string, dstr string) (int, error) {
	if str == "" || dstr == "" {
		return -1, errors.New("empty value")
	}
	pointer := 0
	for i := 0; i < len(str); i++ {
		if str[i] == dstr[pointer] {
			pointer++
		} else {
			pointer = 0
		}
		if pointer == len(dstr) {
			return 0, nil
		}
	}
	return 1, nil
}

func PrimeNumCheck(num int) (int, error) {
	if num < 1 {
		return -1, errors.New("value must be > 0")
	}
	if num == 1 {
		return 1, nil
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return 1, nil
		}
	}
	return 0, nil
}

func ReverceNum(num int32) (int32, error) {
	var rev_num int64
	rev_num = int64(num)
	var res int64
	res = 0
	if rev_num < 0 {
		rev_num *= -1
	}
	for rev_num/10 > 0 || rev_num%10 > 0 {
		curr_num := rev_num % 10
		rev_num /= 10
		res *= 10
		res += curr_num
	}
	if res > 2147483647 {
		res = 0
	}
	if num < 0 {
		res *= -1
	}
	var res_int32 int32
	res_int32 = int32(res)
	return res_int32, nil
}

func RomeNumConvert(num int) (string, error) {
	if num < 1 {
		return "", errors.New("value must be > 0")
	}
	rm := []rome_number{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
	var res string
	for i := 0; i < len(rm); i++ {
		for rm[i].value <= num {
			res += rm[i].name
			num -= rm[i].value
		}
	}
	return res, nil
}
