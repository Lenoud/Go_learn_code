// @FileName : 02goroutine.go
// @Time : 2025/3/22 下午1:25
// @Author : luobozi

package main

import (
	"fmt"
	"time"
)

/*运行main 运行之初就是创建一个协程在执行
一个协程 -- runtime 库
创建一个协程  runtime.g 结构体
*/
// 此处开始编写代码
//func main() {
//	fmt.Println("hello world!")
//
//	//GOMAXPROCS 传入0进去表示查询当前设置的P的数量  传入大于0 就是设置P的数量
//	//fmt.Println(runtime.GOMAXPROCS(0))
//	//fmt.Println(runtime.GOMAXPROCS(8))
//	//fmt.Println(runtime.GOMAXPROCS(0))
//
//	go hello()
//	//默认情况下主协程退出，子协程也会退出
//	time.Sleep(1)
//
//}
//
//func hello() {
//	fmt.Println("hello goroutine")
//}

// 匿名协程
//func main() {
//	go func() {
//		i := 0
//		for {
//			i++
//			fmt.Println("子协程 i=", i)
//			time.Sleep(1 * time.Second)
//		}
//	}()
//	for i := 0; i < 3; i++ {
//		fmt.Println("主协程 i=", i)
//		time.Sleep(1 * time.Second)
//	}
//}

// 协程异常处理
// go语言中如果一个协程崩溃了，所有协程都会退出
func except_test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("程序发生异常！", r)
		}
	}()
	fmt.Println("this is except_test")
	lst := []int{1, 2, 3}
	fmt.Println(lst[3])
}

func except_test2() {
	fmt.Println("this is except_test2")
	time.Sleep(1 * time.Second)
	fmt.Println("this is except_test2 end .....")
}

func main() {
	fmt.Println("this is main")
	go except_test()
	go except_test2()
	time.Sleep(2 * time.Second)
	fmt.Println("end....")

}
