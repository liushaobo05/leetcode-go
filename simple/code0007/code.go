package main

import (
	"fmt"
	"math"
)

func main() {
	x := 123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	var res = 0
	sign := 1
	if x < 0 {
		sign = -1
	}
	x = sign * x

	for x > 0 {
		tmp := x % 10
		res = res*10 + tmp
		x = x / 10
	}

	res = sign * res

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}
