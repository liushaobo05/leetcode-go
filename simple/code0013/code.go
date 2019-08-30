package main

import "fmt"

func main() {
	s := "IX"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	var dic = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var strLen = len(s)
	var res = 0
	var sign = 1
	for i := strLen - 1; i >= 0; i-- {
		tmp := dic[s[i]] * sign
		if s[i] == 'V' || s[i] == 'X' || s[i] == 'C' || s[i] == 'L' || s[i] == 'D' || s[i] == 'M' {
			sign = -1
		}
		res += tmp
	}
	return res
}
