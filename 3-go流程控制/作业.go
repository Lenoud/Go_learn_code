// @FileName : 作业.go
// @Time : 2025/3/13 上午11:07
// @Author : luobozi

package main

import "fmt"

// 此处开始编写代码

// 接收用户名和密码的输入，判断用户名是否为root  密码是否为123456
// 如果用户名和密码输出正确，提示登录成功
// 输入错误，有三次重试机会，三次用完提示登录失败
//func main() {
//	task_1()
//}

func task_1() {

	var i int8 = 0
	for {
		var username string
		fmt.Print("输入用户名：")
		fmt.Scanln(&username)
		var passwd string
		fmt.Print("输入用户名：")
		fmt.Scanln(&passwd)
		if (username == "root") && (passwd == "123456") {
			fmt.Println("登录成功！")
			break
		} else if i >= 2 {
			fmt.Println("登录失败！")
			break
		} else {
			fmt.Println("输入错误，请重试！")
		}
		i++
	}

}
