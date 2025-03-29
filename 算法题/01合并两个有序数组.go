// @FileName : 1.go
// @Time : 2025/3/24 下午2:56
// @Author : luobozi

package main

import "fmt"

//https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
// 此处开始编写代码

/*
为了合并两个有序数组 `nums1` 和 `nums2`，并且保持合并后的数组有序，我们可以利用双指针从后向前遍历的方法。这种方法避免了从前向后遍历时可能出现的覆盖问题，确保较大的元素先放置到正确的位置。
### 方法思路
1. **初始化指针**：使用三个指针，`p1` 指向 `nums1` 的最后一个有效元素（即 `m-1`），`p2` 指向 `nums2` 的最后一个元素（即 `n-1`），`p` 指向合并后数组的末尾（即 `m+n-1`）。
2. **从后向前比较**：比较 `nums1[p1]` 和 `nums2[p2]`，将较大的元素放到 `nums1[p]` 的位置，并将对应的指针向前移动。
3. **处理剩余元素**：如果 `nums2` 中还有元素未处理完，继续将它们复制到 `nums1` 的前面位置。由于 `nums1` 的剩余元素已经处于正确的位置，无需处理。


func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1
	for p2 >= 0 { // nums2 还有要合并的元素
		// 如果 p1 < 0，那么走 else 分支，把 nums2 合并到 nums1 中
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1] // 填入 nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2] // 填入 nums2[p1]
			p2--
		}
		p-- // 下一个要填入的位置
	}
	fmt.Println(nums1)
}


### 代码解释
- **指针初始化**：`p1` 和 `p2` 分别指向两个数组的最后一个有效元素，`p` 指向合并后的数组末尾。
- **循环比较**：当 `p2` 未处理完时，比较当前 `nums1` 和 `nums2` 的元素，将较大的元素放到合并后的位置，并移动相应的指针。
- **处理剩余元素**：如果 `nums2` 中还有元素未处理，继续复制到 `nums1` 的前面位置。由于是从后向前处理，所有元素都能正确放置，无需额外操作。

这种方法确保了时间复杂度为 O(m+n)，空间复杂度为 O(1)，高效地完成了两个有序数组的合并。
*/

func main() {
	m := 3
	n := 4
	nums1 := make([]int, n+m)
	nums1 = []int{1, 2, 3, 0, 0, 0, 0}
	nums2 := make([]int, n)
	nums2 = []int{2, 5, 6, 9}
	merge1(nums1, m, nums2, n)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1 //初始化三个指针
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	fmt.Println(nums1)
}
