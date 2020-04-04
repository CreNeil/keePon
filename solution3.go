package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	bytes := []byte(str)
	for i := 0; i < len(bytes)/2; i++ {
		if bytes[i] != bytes[len(bytes)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(strconv.Itoa(-100))
	fmt.Println(3 / 2)
	fmt.Println(isPalindrome(100))
}
