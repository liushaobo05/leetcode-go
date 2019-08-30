package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	res1 := []int{}
	res := []int{}

	f := filter{}
	f2 := filter{}

	for _, val := range nums1 {
		f.Put(val)
	}

	for _, val := range nums2 {
		if f.Get(val) {
			res1 = append(res1, val)
		}
	}

	for _, val := range res1 {
		if !f2.Get(val) {
			res = append(res, val)
			f2.Put(val)
		}
	}

	return res
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
