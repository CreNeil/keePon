package main

func plusOne(digits []int) []int {
	flag := true
	for i := len(digits) - 1; i >= 0; i-- {
		if flag {
			digits[i] = digits[i] + 1
			if digits[i] >= 10 {
				digits[i] = 0
				flag = true
			} else {
				break
			}
		}
	}
	if digits[0] == 0 {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}

func main() {
	nums := []int{9, 9, 9, 9, 9}
	print(plusOne(nums))
}
