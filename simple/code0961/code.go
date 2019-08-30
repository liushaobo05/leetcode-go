package main

import "fmt"

func main() {
	A := []int{1, 2, 3, 3}
	fmt.Println(repeatedNTimes(A))
}

func repeatedNTimes(A []int) int {
	f := make(filter)
	for _, val := range A {
		if ok := f.Get(val); ok {
			return val
		}

		f.Put(val)
	}

	return 0
}

type filter map[int]bool

// 获取值
func (f filter) Get(key int) bool {
	if _, ok := f[key]; ok {
		return true
	}

	return false
}

// 存入值
func (f filter) Put(key int) {
	f[key] = true
}
