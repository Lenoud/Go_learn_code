// @FileName : 01数组.go
// @Time : 2025/3/13 下午1:13
// @Author : luobozi

package main

import (
	"fmt"
	"strconv"
)

// 此处开始编写代码

// 数组是具有相同唯一类型的一组已经编号且长度固定的数据序列
// 数组的元素可以通过索引来读取或者修改，索引从0开始

// var name [size] type
// var arr [5]int
//var arr = [5]int{1,2,3,4,5}
//arr := [5]int{1:200}  //将下标索引为1的元素初始化为200，其它元素都为默认值

//arr := [...]int{1,2,3,4,5}  //数组长度不确定,可以使用...代替数组的长度，编译器会根据元素的个数自行推断数组的长度

//func main() {
//	arrary()
//
//}

func arrary() {
	//数组定义，初始化
	var arr1 [5]int
	fmt.Println(arr1)

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [5]int{1: 200}
	fmt.Println(arr3)

	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr4)

	//获取数组的长度 --len
	fmt.Println(len(arr4))

	//数组的截取
	sub_arr := arr4[2:4]
	fmt.Println(sub_arr)
	fmt.Printf("%T\n", arr4)
	fmt.Printf("%T\n", sub_arr) //通过截取获取到的数据类型是：切片

	//二维数组
	arr5 := [3][4]int{
		{1, 2, 3},
		{4, 5, 6},
		{7}, //大括号另起一行，这里一定要加,逗号
	}
	fmt.Println(arr5)

	for k, v := range arr4 {
		fmt.Println(k, v)
	}

	//定义两个数组，输出数组里面哪些元素同时在两个数组中
	arr_1 := [...]int{1, 2, 5, 4, 6, 3, 1, 10}
	arr_2 := [...]int{6, 5, 8, 7, 4, 3, 6, 52}
	str_1 := ""
	for _, v := range arr_1 {
		for _, v2 := range arr_2 {
			if v == v2 {
				str_1 += strconv.Itoa(v)
				break
			}
		}
	}
	fmt.Println(str_1)
}
