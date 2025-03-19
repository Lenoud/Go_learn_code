// @FileName : main.go
// @Time : 2025/3/18 下午6:40
// @Author : luobozi

package main

import (
	"bufio"
	"fmt"
	"os"
	"shop_system/login"
	"shop_system/user"
	"strconv"
	"strings"
)

// 此处开始编写代码
/*
使用项目包的结构！
写一个三创购物系统

1.注册
 用户名、密码、账户余额
2.登录
 查看商品信息 map形式，商品信息放在config包
 加入购物车,输入商品编号，购买的数量
 结算，账户余额会减少
3.退出
*/
var login_true bool

func main() {
	for {
		//选择界面
		fmt.Println("\n欢迎来到萝卜超市系统，输入你的选择：")
		fmt.Println("1.注册")
		fmt.Println("2.登录")
		fmt.Println("3.退出")
		fmt.Println("4.注销登录")
		fmt.Print("请输入：")
		select1 := bufio.NewReader(os.Stdin)
		line, _ := select1.ReadString('\n')
		controller, _ := strconv.Atoi(strings.Trim(line, "\n"))

		//选择的处理逻辑
		switch controller {
		case 1: //1.注册
			user1 := map[string]string{
				"name":    "luobozi",
				"passwd":  "1234567890",
				"age":     "18",
				"addr":    "HuNanChangSha",
				"college": "znlkdsw",
				"money":   "5000",
			}
			abc := user.User_regriter(user1)
			if abc == true {
				fmt.Printf("用户%s创建成功！\n", user1["name"])
			} else {
				fmt.Println("用户创建失败!")
			}
		case 2: //2.登录
			var username string
			var passwd string
			fmt.Println("输入用户名不超过10位，密码不少于10位：")
			fmt.Print("输入用户名：")
			fmt.Scan(&username)
			fmt.Print("输入密码：")
			fmt.Scan(&passwd)
			login_true = login.Login(username, passwd)
			if login_true == true {
				login.Show_shop()
			LOGIN_SEL:
				fmt.Println("1.查看商品列表")
				fmt.Println("2.加入购物车")
				fmt.Println("3.结账")
				fmt.Println("4.退出选择")
				fmt.Print("输入你的选择：")
				var select2 uint8
				fmt.Scan(&select2)
				switch select2 {
				case 1:
					login.Show_shop()
					goto LOGIN_SEL
				case 2:
					var select3 string
					fmt.Print("你要购买哪个商品（输入位置编号）：")
					fmt.Scan(&select3)
					login.Shop_cat(select3)
					goto LOGIN_SEL
				case 3:
					login.Overorder()
					goto LOGIN_SEL
				case 4:
					fmt.Println("退出选择")
				default:
					fmt.Println()
				}
			} else {
				fmt.Println("登录失败！")
			}
		case 3: //3.退出程序
			os.Exit(0)
		case 4:
			fmt.Println("注销成功!")
			login.Loginout()
		default:
			fmt.Println("输入3退出程序！")
		}
	}
}
