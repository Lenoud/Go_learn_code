// @FileName : 02指针.go
// @Time : 2025/3/17 下午10:08
// @Author : luobozi

package main

import "fmt"

// 此处开始编写代码
func main() {
	var a int = 100
	var p *int
	p = &a
	fmt.Printf("a变量的地址是：%x\n", &a)
	fmt.Printf("p的值是：%x\n", p)
	fmt.Printf("p的地址是：%x\n", &p)
	fmt.Println("a的值为：", a)
	fmt.Println("解引用指针p得到a值为：", *p)
	*p = 500
	fmt.Println("a的值为：", a)
	fmt.Println("解引用指针p得到a值为：", *p)

	var prt *int
	if prt == nil {
		fmt.Println("prt为空指针")
	}

	//go语言的指针不能参与运算
	//不同类型的指针不允许相互转换
	//不同类型的指针不能比较和相互赋值

}
