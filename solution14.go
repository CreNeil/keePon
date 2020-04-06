package main

import "strings"

func lengthOfLastWord(s string) int {
	strs := strings.Split(s, " ")
	result := 0
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) != 0 {
			result = len(strs[i])
		}
	}
	return result
}

func main() {
	print(lengthOfLastWord("adf 4653"))
}
