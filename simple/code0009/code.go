package main

import "fmt"

func main() {
	num := 0
	fmt.Println(isPalindrome(num))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	revNum := 0
	for num := x; num > 0; {
		tmp := num % 10
		revNum = revNum*10 + tmp
		num = num / 10
	}

	return revNum == x
}
