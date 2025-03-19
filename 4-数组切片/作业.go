// @FileName : 作业.go
// @Time : 2025/3/13 下午8:38
// @Author : luobozi

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 此处开始编写代码
//
//	func main() {
//		//task_1()
//	}
func task_1() {
	//编写一个随机点名系统
	//定义一个切片类型放全班同学的名字
	//按任意键，不限次数随机抽取切片中莫一个同学回答问题
	//按q退出，学生全部点名完了  也退出
	//抽取到的学生就不要再抽了
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Println(rand.Int(), 1*time.Second)
	}
}
