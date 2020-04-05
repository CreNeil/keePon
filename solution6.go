package main

import "strings"

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	keyMap := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	/**
	table[]
	table["(","[","{"]
	"}"
	table["(","["]
	*/
	var table []string
	for i := 0; i < len(s); i++ {
		if len(table) > 0 {
			tmp, ok := keyMap[string(s[i])]
			if ok {
				top := table[len(table)-1]
				if top == tmp {
					table = table[:len(table)-1]
					continue
				}
			}
		}
		table = append(table, string(s[i]))
	}
	// 如果最后栈是空的，说明都匹配上了，返回true
	return len(table) == 0
}

func isValid2(s string) bool {
	for len(s) > 0 {
		if strings.Contains(s, "()") || strings.Contains(s, "[]") || strings.Contains(s, "{}") {
			s = strings.ReplaceAll(s, "()", "")
			s = strings.ReplaceAll(s, "[]", "")
			s = strings.ReplaceAll(s, "{}", "")
		} else {
			return false
		}
	}
	return s == ""
}
func main() {
	println(isValid("((})}"))
	println(isValid2("((})}"))
}
