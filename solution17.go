package main

func mySqrt(x int) int {
	i, j := 1, x
	for i < j {
		mid := (i + j) / 2
		if mid*mid > x {
			if j == mid {
				return j
			}
			j = mid
		} else if mid*mid < x {
			if i == mid {
				return i
			}
			i = mid
		} else {
			return mid
		}
	}
	return x
}

func mySqrt2(x int) int {
	i, j := 1, x
	for i <= j {
		mid := (i + j) / 2
		if mid*mid > x {
			j = mid - 1
		} else if mid*mid < x {
			i = mid + 1
		} else {
			return mid
		}
	}
	return i - 1
}

func main() {
	println(mySqrt2(10))
}
