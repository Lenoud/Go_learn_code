// @FileName : 01流程控制.go
// @Time : 2025/3/13 上午9:51
// @Author : luobozi

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// 此处开始编写代码

// 1.顺序结构
// 2.分支结构
// 3.循环结构
func main() {
	switch_task1()
	switch_task2()
	if_task1()
	for_task1()
}

func if_task1() {
	var str_1 string
	str_1 = "12"
	if v, e := strconv.Atoi(str_1); e != nil {
		fmt.Println("出现错误", e)
	} else if v > 100 {
		fmt.Println(v, "> 100")
	} else {
		fmt.Println(v, "< 100")
	}
}

func switch_task1() {
	select_1 := bufio.NewReader(os.Stdin)
	fmt.Println("输入成绩(1-100)：")
	line, _ := select_1.ReadString('\n')
	number_1, _ := strconv.Atoi(strings.TrimSpace(line))
	switch {
	case number_1 >= 85:
		fmt.Println("成绩等级：A")
	case number_1 >= 60:
		fmt.Println("成绩等级：B")
	default:
		fmt.Println("成绩等级：C")
	}
}

func switch_task2() {
	var select_1 int
	fmt.Println("请输入你的选择（1,2）：")
	fmt.Scanln(&select_1)

	switch select_1 {
	case 1:
		fmt.Println("选择了登录")
	case 2:
		fmt.Println("选择了注册")
	default:
		fmt.Println("输入有误！")
	}

}

func for_task1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var x int8 = 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	var choice string
	for choice != "quit" {
		fmt.Print("输入选择(quit退出选择):")
		fmt.Scanln(&choice)
	}

	for k, v := range "abcdefg" {
		fmt.Println(k, string(v))
	}

	//for {
	//	//无限死循环  只能通过条件+break的方式退出
	//	fmt.Println("无限死循环")
	//}

	//break退出当前循环  continue退出本次循环
	//go语言for循环有自己的局部变量空间
	//在for循环中创建的变量只能在循环体中使用
	//for循环中创建的变量和父作用域的同名变量互不不干涉
	a := 10
	fmt.Printf("%d addrees is %d\n", a, &a)
	for _, v := range "abc" {
		fmt.Println("当前值为：", string(v))
		a := 100
		fmt.Printf("%d ,addrees is %d\n", a, &a)
	}

	//goto 语句 慎用！
	num1 := 1
ABC:
	fmt.Println("goto代码块输出：", num1)
	for num1 < 10 {
		num1++
		if num1 == 5 {
			goto ABC
		}
	}
	fmt.Printf("%T", time.Now().Year())
	fmt.Println()
	fmt.Println()

}
