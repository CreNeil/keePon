package main

import "fmt"

func romanToInt(s string) int {
	var result int
	temp := 0
	for i := len(s) - 1; i >= 0; i-- {
		num := getNumber(string(s[i]))
		if num < temp {
			result -= num
		} else {
			result += num
		}
		temp = num
	}
	return result
}

func getNumber(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

func main() {
	fmt.Println(romanToInt("LVIII"))
}
