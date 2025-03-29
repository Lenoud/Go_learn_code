package main

import "fmt"

func removeDuplicates(nums []int) int {
	p1 := 0 // 维护可插入位置
	for _, num := range nums {
		// 允许最多两个相同元素
		if p1 < 2 || num != nums[p1-2] {
			nums[p1] = num
			p1++
		}
	}

	fmt.Println(nums[:p1]) // 输出修改后的数组
	return p1
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	length := removeDuplicates(nums)
	fmt.Println("New Length:", length)
}
