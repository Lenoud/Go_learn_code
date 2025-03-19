// @FileName : 作业.go
// @Time : 2025/3/12 下午5:19
// @Author : luobozi

package main

import (
	"fmt"
	"strconv"
)

// 此处开始编写代码
//1.从键盘接收一个字符串，判断这个字符串是否为回文
//2.字符串压缩”aabbcccddd“ -> ”a2b2c3d3“

func main() {
	task_1()
	//task_2()
	task_3()
}

func task_1() {
	str_1 := "上海自来水来自海上"
	//fmt.Scanln(&str_1)
	str_2 := ""

	for _, v := range str_1 {
		str_2 = string(v) + str_2
	}

	if str_1 == str_2 {
		fmt.Println(str_1, "是一个回文")
	} else {
		fmt.Println(str_1, "不是一个回文")
	}

}

// 2.字符串压缩”aabbcccddd“ -> ”a2b2c3d3“
func task_2() {
	str_1 := "aaaabccbcccccdddddabcgggggg"
	count := 0
	str_2 := ""
	str_3 := ""
	for _, v := range str_1 {
		if str_2 == string(v) {
			count += 1
		} else {
			str_2 = string(v)
			count = 1
		}
		str_3 = str_2 + strconv.Itoa(count)
		fmt.Println(count, str_3)

	}

}

// deepseek
func task_3() {
	str_1 := "aaabbbbcccca"
	var preChar string // 用于记录前一个字符
	count := 0         // 计数器
	result := ""       // 最终结果字符串

	for _, v := range str_1 {
		nowChar := string(v)
		if nowChar == preChar {
			count++ // 相同字符，计数器加1
		} else {
			// 当字符变化且前一个字符存在时，将前一个字符及其计数添加到结果
			if preChar != "" {
				result += preChar + strconv.Itoa(count)
			}
			preChar = nowChar // 更新为当前新字符
			count = 1         // 重置计数器
		}
	}

	// 处理最后一个字符的计数
	if preChar != "" {
		result += preChar + strconv.Itoa(count)
	}

	fmt.Println("压缩结果:", result)
}
