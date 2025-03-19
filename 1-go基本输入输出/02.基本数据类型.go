// @FileName : 02.基本数据类型.go
// @Time :2025/2/10 14:53
// @Author : luobozi

package main

import (
	"fmt"
	"unsafe"
)

// 有符号整形
// int8 int16 int32 int64  int
// 默认int 就是看操作系统
// 无符号整形
// uint8 uint16 uint32 uint64 uint
// 默认uint 也是看操作系统
// unitprt 无符号整数类型指针

var a int8  //-127 -- 127
var b int16 //
var x uint8
var y uint8

//func main() {
//	int_1()
//	fmt.Println("------------------")
//	float_1()
//	fmt.Println("------------------")
//	complex_1()
//}

func int_1() {

	a = 127 //不能赋值超过127
	b = 200
	fmt.Println(a, b)

	//无符号位  注意计算溢出  就是如果计算结果为负数 会导致计算溢出
	x = 1
	y = 2
	fmt.Println("无符号计算溢出：1-2 =", x-y) //这种情况下会输出255 和计算结果不符，就是溢出了

	//int类型看系统位数
	//unsafe.Sizeof返回字节大小    8字节= 64位  意味着我们操作系统是64位的
	var x int
	x = 100
	fmt.Println("我们操作系统的int占用8字节=8位*8=64位操作系统，xByte=", unsafe.Sizeof(x))

	//字符串默认就占用16个字节
	var str1 string
	fmt.Println("字符串默认就占用16个字节: str1 =", unsafe.Sizeof(str1))

	//进制表示
	var bin_num1 = 0b11 //二进制 0 and 1
	var oct_num1 = 0o11 //八进制 1-7
	var hex_num1 = 0x11 //十六进制 1-9 a-f
	fmt.Println("0b11二进制：", bin_num1, "\n0o11八进制：", oct_num1, "\n0x11十六进制：", hex_num1)
	fmt.Printf("%T", bin_num1)

	// int64 和 int 不能直接进行计算，即使他们都是64位的——> go语言严格区分数据类型
	var type_64 int64
	var type_int int
	fmt.Println("int64 和 int 不能直接进行计算", type_int, type_64)
	//fmt.Println(type_int + type_64)
	//invalid operation: type_int + type_64 (mismatched types int and int64)

}

func float_1() {
	// 浮点型： float32 float 64  ——> 是不精确的 是因为浮点数小数换算成二进制的机制导致的
	//float32  三个部分组成  1符号   8指数位 23小数位
	//float64			   1符号   11指数位 52小数位

	fmt.Println("浮点型： float32 float 64  ——> 是不精确的 是因为浮点数小数换算成二进制的机制导致的")
	var v1 float32
	var v2 float32
	var v3 float32
	v1 = 0.111111
	v2 = 0.222222
	v3 = 0.333333
	if v1+v2 == v3 {
		fmt.Println("0.111111 + 0.222222 == 0.333333 在go中： 相等")
	} else {
		fmt.Println(" 0.111111 + 0.222222 == 0.333333 在go中： 不相等")
	}

	//如果对于小数点要求比较高的场景使用第三方库 decimal
	fmt.Println("如果对于小数点要求比较高的场景使用第三方库 decimal")

}

func complex_1() {
	//复数 complex64 complex128
	var z1 complex64 = 5 + 4i
	fmt.Println(z1)
}
