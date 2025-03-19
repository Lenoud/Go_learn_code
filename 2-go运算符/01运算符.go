// @FileName : 01运算符.go
// @Time : 2025/3/12 下午7:54
// @Author : luobozi

package main

import (
	"fmt"
	"math"
)

// 此处开始编写代码

func main() {
	go_math()
	relational()
}

// 算数运算符 + - * / % ++ --
func go_math() {
	x, y := 8, 5
	fmt.Printf("x = %d ,y = %d\n", x, y)
	fmt.Println("x/y除法运算:", x/y) //整数运算只能得到整数部分
	fmt.Println("float64(x)/float64(y)浮点运算:", float64(x)/float64(y))
	fmt.Println("x%y取余运算:", x%y)
	x++
	y--
	fmt.Println("x++自增运算:", x)
	fmt.Println("y--自减运算:", y)

	//除了上面的基本运算外，还有一个math库提供了许多数学函数和常量
	fmt.Println("除了上面的基本运算外，还有一个math库提供了许多数学函数和常量")
	fmt.Println("math.Pi=", math.Pi)
	fmt.Println("math.Pow幂运算", math.Pow(5, 3))
	fmt.Println("math.Abs取绝对值：", math.Abs(-100))
}

// 关系运算符号 < > = != => <=
// 逻辑运算符  -- 注意短路运算
// 将多个条件组合在一起去判断 进行条件的与、或、非
// &&  x && y  x为真，整个表达式看y的结果 ， 如果x为假， 不会去判断y的结果了（称之为短路y表达式的计算）
// ||  x || y  x为真，整个表达式结果为真，短路y的表达式
// !
func relational() {
	//关系运算符，不能连用  不能 5 < x < 15
	x := 10

	if 5 > x && 15 > x {
		fmt.Println(x)
	}

	//位运算符
	fmt.Println(5 | 6)  //或运算
	fmt.Println(5 & 6)  //与运算
	fmt.Println(5 ^ 6)  //异或运算
	fmt.Println(5 << 2) //左移运算
	fmt.Println(5 >> 2) //右移运算

	//赋值运算符  = += -= /=

	//其他运算符号  & * -> <-
}
