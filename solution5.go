package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var result string
	if len(strs) == 0 {
		return result
	}
	if len(strs) < 2 {
		return strs[0]
	}
	result = strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(result) {
			result = strs[i]
		}
	}
	var j = len(result)
LOOP:
	for j >= 0 {
		for k := 0; k < len(strs); k++ {
			if !compare(result[:j], strs[k]) {
				j--
				goto LOOP
			}
		}
		return result[:j]
	}
	return result
}

func compare(str1 string, str2 string) bool {
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}

func main() {
	//strs := []string{"flower","flow","flight"}
	strs := []string{"ca", "a"}
	fmt.Println(longestCommonPrefix(strs))
}
