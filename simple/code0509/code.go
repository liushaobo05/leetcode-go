package main

import "fmt"

func main() {
	N := 10
	fmt.Println(fib(N))
}

func fib(N int) int {
	if N == 1 {
		return 1
	}

	if N == 0 {
		return 0
	}

	return fib(N-1) + fib(N-2)
}
