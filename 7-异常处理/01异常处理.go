// @FileName : 01异常处理.go
// @Time : 2025/3/18 下午9:43
// @Author : luobozi

package main

import (
	"errors"
	"fmt"
	"os"
)

// 此处开始编写代码
//go语言中尽量避免抛出异常
/*
一般发生的异常错误，都期望程序员自己处理
在函数的实现过程中，出现了不能处理的错误，可以返回给调用者处理
*/

func main() {
	//if e := hello("luobozi"); e != nil {
	//	fmt.Println(e)
	//}
	//fmt.Println(getelem(3))
	//arr := getelem_2(3)
	arr := getelem_2(3)
	fmt.Println("继续执行：", arr)

}

// 自己产生错误类型
// 内置类型 error
// 产生错误的内置包 errors
// 自定义的error 往往都是能够预知的
// 不能预知 go -- panic --程序终止了
func hello(name string) error {
	if len(name) == 0 {
		return errors.New("NameError: name is empty!!")
	}
	fmt.Println("hello !", name)
	return nil

}

func getelem(index int) (int, error) {
	arr := [3]int{10, 20, 30}
	//return arr[index]
	//尽量避免panic
	if -1 < index && index < len(arr) {
		return arr[index], nil
	}
	return 0, errors.New("index out of range")
}

func err_1() {
	_, err := os.OpenFile("test.txt", os.O_RDONLY, 777)
	if err != nil {
		fmt.Println("err：打开文件失败")
	}
}

func recover_1() {
	//recover() 函数 用于终止异常处理流程 类似try ... except ...
	//recover() 要配置defer执行
	//defer 延迟处理语句 先定义后执行
	fmt.Println("this is defer test!")
	defer fmt.Println(100)
	defer fmt.Println(200)
	fmt.Println("this is defer test 2...")
}

func getelem_2(index int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("something error!", r)
		}
		fmt.Println("defer语句！")
	}() //不管有没有panic 在退出函数执行之前 defer 语句都会执行
	arr := [3]int{5, 6, 7}
	panic("手动抛出异常") //类似于python中的raise
	return arr[index]
}

func getelem_3(index int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("something error!", r)
		}
		fmt.Println("defer语句！")
	}() //不管有没有panic 在退出函数执行之前 defer 语句都会执行
	arr := [3]int{5, 6, 7}
	return arr[index]
}
