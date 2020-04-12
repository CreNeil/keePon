package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	var length int
	gap := len(b) - len(a)
	if gap > 0 {
		length = len(b)
		for i := 0; i < gap; i++ {
			a = "0" + a
		}
	} else if gap < 0 {
		length = len(a)
		for i := 0; i < -gap; i++ {
			b = "0" + b
		}
	} else {
		length = len(a)
	}
	result := ""
	flag := false
	res := ""
	for i := length - 1; i >= 0; i-- {
		res, flag = judge(string(a[i]), string(b[i]), flag)
		result = res + result
	}
	if flag {
		result = "1" + result
	}
	return result
}
func judge(a string, b string, f bool) (string, bool) {
	ai, _ := strconv.Atoi(a)
	bi, _ := strconv.Atoi(b)
	pos := 0
	if f {
		pos = 1
	}
	switch ai + bi + pos {
	case 0:
		return "0", false
	case 1:
		return "1", false
	case 2:
		return "0", true
	case 3:
		return "1", true
	}
	return "0", false
}
func main() {
	a := "10001"
	b := "10100"
	println(addBinary(a, b))
}
