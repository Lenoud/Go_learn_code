// @FileName : 作业.go
// @Time : 2025/3/16 下午8:32
// @Author : luobozi

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"unicode/utf8"
)

// 此处开始编写代码

//1注册  --固化，写入磁盘 username passwd age addr
//2登录	-- username passwd
//3退出
//无限次数，将注册登录封装成函数

//func main() {
//
//	var cheio int8
//	for {
//		fmt.Println("1.注册")
//		fmt.Println("2.登录")
//		fmt.Println("3.退出")
//		fmt.Print("请输入数字：")
//		fmt.Scanln(&cheio)
//		switch cheio {
//		case 1:
//			user, e := enroll("luobozi", "123456789", "12", "cs")
//			if e == false {
//				fmt.Println("注册失败！")
//			} else {
//				fmt.Println("注册成功", user)
//			}
//		case 2:
//
//		case 3:
//			os.Exit(0)
//		}
//	}
//
//}

func login(username, password string) bool {

	return false
}

func enroll(username, passwd, age, addr string) (map[string]string, bool) {
	usr_lan := utf8.RuneCountInString(username)
	psd_lan := utf8.RuneCountInString(passwd)

	userinfo := make(map[string]string, 100)

	if usr_lan <= 9 || psd_lan <= 8 {
		userinfo["username"] = username
		userinfo["passwd"] = passwd
		userinfo["age"] = age
		userinfo["addr"] = addr
		user_str, e := json.Marshal(userinfo)

		//固化数据
		if e != nil {
			fmt.Println("转换出错")
		} else {
			fmt.Println("用户信息转换成json:", string(user_str))
			err := os.WriteFile("D:\\Environment\\GO\\gopath\\src\\5-函数\\person.json", user_str, 0644)
			if err != nil {
				fmt.Println("保存错误", err)
			}
		}
		return userinfo, true

	} else {
		fmt.Println("用户名控制在9位数或以下！")
		fmt.Println("密码需要8位数或以上！")
		return nil, false
	}
	return nil, false
}
