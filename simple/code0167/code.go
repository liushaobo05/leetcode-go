package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(numbers []int, target int) []int {
	var res = []int{}

	f := make(filter)

	for idx, val := range numbers {
		if v, ok := f.Get(val); ok {

			res = append(res, v+1, idx+1)
			return res
		}

		key := target - val
		f.Put(key, idx)
	}

	return res
}

type filter map[int]int

// 获取值
func (f filter) Get(key int) (int, bool) {
	if val, ok := f[key]; ok {
		return val, true
	}

	return -1, false
}

// 存入值
func (f filter) Put(key int, val int) {
	f[key] = val
}
