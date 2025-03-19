// @FileName : 作业.go
// @Time : 2025/3/12 下午8:45
// @Author : luobozi

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// 此处开始编写代码

//1.从键盘接受用户输入的身份证号码 以4开头18位数字，判断输入合法性
//获取身份证上的出生年月日，判断这个人是否大于18岁
//大于18岁提示录用，否则提示劝退

//func main() {
//	task_1()
//}

func task_1() {
	fmt.Println("请输入4开头的18位身份证号：")
	read := bufio.NewReader(os.Stdin)
	line, _ := read.ReadString('\n')
	//431382200402060216
	//431382200705050216
	//431382200802060216
	//831382200802060216
	//831382200802vc0216
	//清洗换行符号
	number1_str := strings.Trim(line, "\n")

	// 转换数字
	if _, e := strconv.Atoi(number1_str); e != nil { //判断输入是否是纯数字
		fmt.Println("请输入数字！")
	} else if lens := int8(utf8.RuneCountInString(number1_str)); lens != 18 { // 获取长度是否为18
		fmt.Println("输入号码不为18位数！")
	} else if !strings.HasPrefix(line, "4") { //获取开头字符是否为4
		fmt.Println("输入号码不以4开头！")
	} else {
		//获取身份证上的出生日期
		year, _ := strconv.Atoi(line[6:10])
		month, _ := strconv.Atoi(line[10:12])
		day, _ := strconv.Atoi(line[12:14])

		//获取当前时间
		nyear := time.Now().Year()
		nmonth := int(time.Now().Month())
		nday := time.Now().Day()

		age := nyear - year
		//判断今年有没有过生日，如果没有，年龄就减一
		if month > nmonth || (nmonth == month && nday < day) {
			age--
		}

		fmt.Println("当前年龄：", age, "出生年月：", year, month, day, "当前时间：", nyear, nmonth, nday)
		if age >= 18 {
			fmt.Println("录用")
		} else {
			fmt.Println("劝退")
		}
	}
}
