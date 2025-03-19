// @FileName : 01.print_and_input.go
// @Time :2025/2/10 14:09
// @Author : luobozi

package main

import (
	"bufio"
	"fmt"
	"os"
)

//func main() {
//	//使用fmt包 输入 fmt.Scan fmt.Scanf fmt.Scanln
//	//var a int
//	//var b int
//	//fmt.Println("请输入a和b")
//	//& get address &a get var a of mem address
//	//fmt.Scan(&a, &b)
//	//fmt.Println(a, b)
//	test()
//
//}

func test() {
	//换行结束输入
	//var a int
	//var b int
	//fmt.Println("请输入a和b")
	//fmt.Scanln(&a, &b)
	//fmt.Println("a is:", a, "b is:", b)

	//严格按照格式输入
	var a int
	var b int
	fmt.Println("请输入a和b")
	fmt.Scanf("a=%d,b=%d", &a, &b)
	fmt.Println("a is:", a, "b is:", b)

	//使用bufio进行输入
	//创建一个从键盘接受数据的对象
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("开始从键盘输入：")
	//
	line, e := reader.ReadString('\n')
	fmt.Println("line is :", line)
	fmt.Println("e is :", e)

	//使用匿名变量接受 _ 接受
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("开始从键盘输入：")
	////
	//line, _ := reader.ReadString('\n')
	//fmt.Println("line is :", line)

	//输出print printf println
	//fmt.Print("hello")
	//fmt.Print("world")

	// 自动添加换行符
	//fmt.Println("hello world!")

	//指定格式
	//name := "luobozi"
	//age := 21
	//fmt.Printf("myname is %s , age is %d", name, age)

	//for i := 1; i < 10; i++ {
	//	for j := 1; j <= i; j++ {
	//		number := i * j
	//		fmt.Printf("%d * %d = %d    ", i, j, number)
	//	}
	//	fmt.Println()
	//}

}
