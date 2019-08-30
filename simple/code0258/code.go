package main

import "fmt"

func main() {
	var num = 119
	fmt.Println(addDigits(num))
}

func addDigits(num int) int {
	res := digits(num)
	for res >= 10 {
		res = digits(res)
	}

	return res
}

func digits(num int) int {
	res := 0
	for num > 0 {
		tmp := num % 10
		num = num / 10
		res += tmp
	}

	return res
}
