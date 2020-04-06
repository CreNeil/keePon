package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	var result strings.Builder
	if n == 1 {
		return "1"
	}
	content := countAndSay(n - 1)
	count := 1
	for i := range content[:len(content)-1] {
		if content[i] != content[i+1] {
			result.WriteString(strconv.Itoa(count))
			result.WriteByte(content[i])
			count = 1
		} else {
			count++
		}
	}
	result.WriteString(strconv.Itoa(count))
	result.WriteByte(content[len(content)-1])
	return result.String()
}
