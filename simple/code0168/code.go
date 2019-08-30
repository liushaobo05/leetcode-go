package main

import "fmt"

func main() {
	n := 27
	fmt.Println(convertToTitle(n))
}

func convertToTitle(n int) string {
	res := []byte{}
	for n > 0 {
		res = append(res, byte('A'+(n-1)%26))
		n = (n - 1) / 26
	}

	return string(res)
}
