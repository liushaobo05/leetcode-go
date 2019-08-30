package main

import "fmt"

func main() {
	deck := []int{r1, 1, 1, 2, 2, 2, 3, 3}
	fmt.Println(hasGroupsSizeX(deck))
}

func hasGroupsSizeX(deck []int) bool {
	var res = 0
	for _, val := range deck {
		res ^= val
	}

	return res == 0
}
