// @FileName : main.go
// @Time :2025/2/10 13:04
// @Author : luobozi

package main

import (
	"fmt"
)

func main() {
	relational()
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
