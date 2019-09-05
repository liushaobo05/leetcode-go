package main

import "fmt"

func main() {
	s := "B"
	fmt.Println(titleToNumber(s))
}

func titleToNumber(s string) int {
	var res = 0
	for i := 0; i < len(s); i++ {
		tmp := s[i] - 'A' + 1
		res = res*26 + int(tmp)
	}

	return res
}
