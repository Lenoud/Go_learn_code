// @FileName : 移除元素.go
// @Time : 2025/3/24 下午3:59
// @Author : luobozi

package main

import "fmt"

/*
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	fmt.Println(nums)
	return left
}
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。
假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：
更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
返回 k。
*/
// 此处开始编写代码
func main() {
	a := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(a)
}

func removeElement(nums []int, val int) int {
	e1 := len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...) //移除nums[i]这个元素
			nums = append(nums, 0)                 //为了使得切片大小不变，往后追加一个元素
			e1--
		}
	}
	fmt.Println(nums)
	return e1
}
