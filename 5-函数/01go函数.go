// @FileName : 01go函数.go
// @Time : 2025/3/15 下午8:40
// @Author : luobozi

package main

import (
	"fmt"
	"unicode/utf8"
)

// 此处开始编写代码
//go是函数式编程范式的语言，由函数组织结构
//可以把函数作为参数传递，也能返回函数
//支持多个返回值，支持命名参数返回，支持可变长参数，支持匿名参数，支持闭包，可以把函数作为参数传递、赋值、返回
//不支持有名函数的嵌套，不支持重载，不支持默认参数

func main() {
	test1()
	fmt.Println(test2(5, 6))
	x, y := test3(5, 6) //多个返回值
	fmt.Println(x, y)
	z := test4(5, 6)
	fmt.Println(z)

	//检查密码
	var pass string
	fmt.Println("输入密码：")
	fmt.Scanln(&pass)
	check := pass_check(pass)
	if check == true {
		fmt.Println("符合要求")
	} else {
		fmt.Println("不符合要求")
	}

	//调用可变长参数
	fmt.Println("调用可变长参数")
	result1 := test5()
	result2 := test5(1)
	result3 := test5(1, 2)
	result4 := test5(1, 2, 3)
	fmt.Println(result1, result2, result3, result4)

	//匿名函数
	noname := func() {
		fmt.Println("noname是一个匿名函数")
	}
	noname()
	//定义匿名函数时调用
	usefunc := func(num1, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println(usefunc)

	//函数作为返回值
	myf := myfunc()
	result := myf()
	fmt.Println(result)

	//把函数作为参数传递
	result5 := myfunc3(test4, 10, 20)
	fmt.Println(result5)

	// 自定义类型
	//type a int //自定义一个a类型，底层使用的还是int类型
	//var x1 a = 123
	//var y1 int = 456
	//fmt.Println(x1 + y1)
	//还是不可以计算
	//invalid operation: x1 + y1 (mismatched types a and int)

}

func test1() {
	fmt.Println("this is test1")
}

// 接收参数，返回数据
func test2(a, b uint8) uint8 {
	return a + b //return返回 函数就退出了

}

// 支持多个返回值
func test3(a int, b int) (int, int) {
	result1 := a + b
	result2 := a - b
	return result1, result2
}

// 支持命名参数返回
func test4(a, b int) (result int) {
	result = a * b
	return
}

//从键盘接收一串密码的输入，判断它是否符合要求
//长度在8-18之间，密码需要大写小写数字，返回真假

func pass_check(password string) (pass bool) {
	pass_len := utf8.RuneCountInString(password)
	fmt.Println(pass_len)
	if pass_len < 8 || pass_len > 18 {
		return false
	}
	var lowercheck, uppercheck, numbercheck bool
	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			lowercheck = true
		} else if char >= 'A' && char <= 'Z' {
			uppercheck = true
		} else if char >= '0' && char <= '9' {
			numbercheck = true
		}
	}
	return lowercheck && uppercheck && numbercheck
}

// 支持可变长参数
func test5(nums ...int) int {
	fmt.Printf("%T --> %v\n", nums, nums)
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

//packing(打包) 和unpacking(解包)
//函数传参 packing test(1,2,3) --> test5(nums ... int)
//					1,2,3 -->打包成 切片[1 2 3] 传递给--> nums
//如果点放在切片后面 unpacking [6 7 8]... 解包成-->6,7,8

// 不支持有名函数嵌套
//# command-line-arguments
//.\01go函数.go:109:7: syntax error: unexpected test7, expecting (
//.\01go函数.go:113:1: syntax error: non-declaration statement outside function body
//func test6() {
//	fmt.Println("this is test6")
//	func test7() {
//		fmt.Println("this is test6")
//	}
//}

// 不支持重载 一个包中函数的名字不能重复！
//.\01go函数.go:118:6: test5 redeclared in this block
//.\01go函数.go:92:6: other declaration of test5
//func test5() {
//	fmt.Println("this is test5")
//}

//不支持默认参数
//func test6(a int = 10)  {
//	fmt.Printf("%T --> %v\n", a, a)
//
//}

// 将函数作为返回值
func myfunc() func() int {
	return func() int {
		fmt.Println("这是返回的函数")
		fmt.Println("返回一个整形")
		return 10
	}
}

//func myfunc2(num1,num2 int) func (num1 int,str1 string) int {}

func myfunc3(f func(int, int) int, num1, num2 int) int {
	return f(num1, num2)
}

// 当函数太长，使用这种方式将函数作为参数传入
type myFunc func(int, int) int

func myfunc4(f myFunc, num1, num2 int) int {
	return f(num1, num2)
}

// 闭包
//一般情况下函数内的局部变量，调用的时候生成，而结束的时候会释放

// 但是在闭包条件下，局部变量不会随着调用的结束而释放
// 1、有内外函数
// 2、外函数要返回内涵数
// 3、内函数要引用外函数的变量
func outer(x int) func(int) int {
	fmt.Println("x is :", x)
	return func(y int) int {
		fmt.Println("inner x is :", x)
		fmt.Println("inner x is :", y)
		return x + y
	}
}
