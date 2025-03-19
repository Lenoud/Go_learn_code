// @FileName : 03.字符串.go
// @Time :2025/2/10 15:44
// @Author : luobozi

package main

//字符类型 byte rune

//编码 人类语言---机器语言  映射
//ascii -- 英文字符
//unicode -- 万国码 utf8 gbk 是unicode具体实现，实现存入磁盘,它们都有不同的标志位来区分

//byte -- uint8   英文字符使用byte
//runne -- uint32  中文等其它字符使用runne

//字符串类型 string 双引号包裹
//如果在go里面使用字符串，一定要使用双引号！

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

//func main() {
//	string_1()
//	string_2()
//}

func string_1() {

	var c1 byte
	c1 = 97
	fmt.Println(c1, "\nsrting 将字符类型转换成文字显示:", string(c1)) //

	var c2 = '中'                      //单引号得到的是字节码,不能用单引号引起多个字符
	fmt.Println("单引号输出中的字节码编码是：", c2) //这里输出的就是中的字节码，而不是 字符
	fmt.Printf("%T\n", c2)

	//如下情况会报错《more than one character in rune literal》
	//var c3 = '中国'

	//双引号原样输出
	var str1 string
	str1 = "中" //双引号包裹字符串
	fmt.Println("双引号输出字符，而不是字符码：", str1)
	var str2 string
	str2 = "中国"
	fmt.Println(str2)

	//统计字符串的长度 -- len(s1 srting类型) -- utf8.RuneCountInString(s1 srting类型)
	var str3 string
	var str4 string
	str3 = "abc萝卜子"
	//len 统计字节数，不是字符数
	fmt.Println("str3=abc萝卜子，占用的字节数是（中文占用3字节）：", len(str3))
	str4 = "...中国china"
	//utf8.RuneCountInString(s1 srting类型) 统计字符串的长度 --> 统计unicode码的个数
	count := utf8.RuneCountInString(str4)
	fmt.Println("str4=...中国china，的字符个数是：", count)

	//转义字符  --  含有特殊含义的字符
	// \t制表符  \n换行符 \" 特殊符号本身
	fmt.Println("--------------------\n\n\n转义字符：")
	var str5 string
	var str6 string
	str5 = "hello\n\"\t\\world"
	//反引号包裹的字符串——> 原样输出
	str6 = `反引号包裹的字符串——> 原样输出: hello \n\t\\\" world`
	fmt.Println(str5)
	fmt.Println(str6)

	//字符串的截取 [start:end] 按字节截取
	fmt.Println("--------------------\n\n\n")
	fmt.Println("字符串的截取 [start:end] 按字节截取")
	var str7 string
	var str8 string
	str7 = "hello world!"
	str8 = "你好世界" //反引号
	/*
		hello wor l  d
		123456789 10 11
		[) 左闭右开区间
	*/
	fmt.Println(str7[0:4])
	fmt.Println("省略开始：", str7[:4])
	fmt.Println("省略结束：", str7[4:])
	fmt.Println("中文不太支持截取操作（按字节截取，中文3字节）：", str8[0:8])

	//字符串和其它数据的转换 -- strconv 库
	//将整形转成字符串
	num1 := 200
	str9 := strconv.Itoa(num1)
	fmt.Printf("将整形转成字符串： \nstr9 is %s,type is %T\n", str9, str9)

	//字符串转换成整形
	str10 := "100"                 //必须要数字的字符串才能转换，abc不能转成整形！
	num2, e := strconv.Atoi(str10) //返回两个值  不能转换e就会拿到错误信息
	if e != nil {                  //nil 为空表示 没有错误信息
		fmt.Println(e)
	} else {
		fmt.Printf("字符串转成整形： \nstr9 is %d,type is %T\n", num2, num2)
	}
	/*
		   % -- %s 字符串
				%c 转换成字符的形式显示 类似与string()
				%d 十进制
				%x 十六进制
				%% %号本身
				%T 输出类型
				%.2f保留两位小数
				%v 输出具体值
	*/

	fmt.Println("--------------------\n")
	//将格式化好的字符串赋值给另一个变量
	str11 := fmt.Sprintf("my name is %s,age is %d", "luobozi", 21)
	fmt.Println("格式化赋值：", str11)
	fmt.Println("--------------------\n")
	//字符串的遍历
	fmt.Println("字符串的遍历")
	str12 := "abcdefg"
	str13 := "hello,萝卜子"

	//按照字节遍历
	fmt.Println("按照字节遍历")
	for i := 0; i < len(str12); i++ {
		fmt.Printf("%v(%c) type: %T \n", str12[i], str12[i], str12[i])
	}

	//按照字符个数遍历
	fmt.Println("\n按照字符个数遍历")
	for _, i := range str13 { //连个值，一个下标，一个具体值，这里下标使用匿名函数接受，因为使用不到
		fmt.Println(i, string(i))
	}

}

func string_2() {
	//接收用户从键盘输入的任意字符串，统计其中英文字母、数字、其他字符的个数

	//统计输入的字符串中，出现的最多的字符和个数，最少字符和个数

	//var input string
	//fmt.Println("请输入一串字符:")
	//fmt.Scanln(&input)
	//fmt.Println(input)
	//max := 0
	//max_char := ""
	//min := utf8.RuneCountInString(input) + 1
	//min_char := ""
	//for _, str_num := range input {
	//	count := 0
	//	for _, str_num_2 := range input {
	//		if str_num_2 == str_num {
	//			count += 1
	//		}
	//	}
	//	if count > max {
	//		max = count
	//		max_char = string(str_num)
	//	}
	//	if count < min {
	//		min = count
	//		min_char = string(str_num)
	//	}
	//}
	//fmt.Println(max, max_char)
	//fmt.Println(min, min_char)

	fmt.Println("--------------------\n")
	//字符串的常用方法 -- strings 包
	//字符串的查找和替换 Contains(s1,s2 string类型) 返回bool 判断字符串s1中是不是包含s2
	//统计字符出现的次数Count(s1, s2 string类型)  计算s1中出现s2的次数
	//Idenx(s1,s2 string类型) 返回s2在s1中第一次出现的位置，若不存在就返回-1
	//Replace(s1,old,new string类型,int 类型)  将s1中前n个old子串替换成new n为-1 表示全部替换
	//Split(s1,sep string类型)根据sep将s1分割成几个部分 返回的是一个数组
	//Tolower(s1 string类型)将s1转换成小写
	//ToUpper(s1 string类型)将s1转换成大写  用在大小写不敏感的场景中
	//Trim(s1,s2 string类型) 将s1两侧的s2字符去除
	//HasPrefix(s1,s2 string类型) 判断s1是不是以s2开头
	//HasSuffix(s1,s2 string类型) 判断s1是不是以s2结尾

	fmt.Println("字符串的常用方法 -- strings 包")
	//字符串的常用方法 -- strings 包
	//字符串的查找和替换 Contains(s1,s2 string类型) 返回bool 判断字符串s1中是不是包含s2
	str_2 := "abcdeffsadfasdf"
	if strings.Contains(str_2, "ab") { //判断子串字符（substr）是否在字符串中
		fmt.Println("条件成立，子串在str_2中")
	}
	//统计字符出现的次数Count(s1, s2 string类型)  计算s1中出现s2的次数
	fmt.Println("a在str_2中出现的次数：", strings.Count(str_2, "a"))

	//Idenx(s1,s2 string类型) 返回s2在s1中第一次出现的位置，若不存在就返回-1
	en_str_2 := "1234567"
	fmt.Println("子串 1 出现在en_str_2下标：", strings.Index(en_str_2, "1"))
	fmt.Println("子串 4 出现在en_str_2下标：", strings.Index(en_str_2, "4"))

	//Replace(s1,old,new string类型,int 类型)  将s1中前n个old子串替换成new n为-1 表示全部替换
	en_str_3 := "hello,world,luobozi,world"
	fmt.Println("替换匹配到的第一个：", strings.Replace(en_str_3, "world", "HuNanChangSha", 1))
	fmt.Println("替换全部匹配成功的字符串：", strings.Replace(en_str_3, "world", "HuNanChangSha", -1))

	//Split(s1,sep string类型)根据sep将s1分割成几个部分 返回的是一个数组
	en_str_4 := "hello/world/HuNan/ChangSha/luobozi/"
	fmt.Println("en_str_4通过 / 符号分割后返回一个数组：", strings.Split(en_str_4, "/"))

	//ToLower(s1 string类型)将s1转换成小写
	//ToUpper(s1 string类型)将s1转换成大写  用在大小写不敏感的场景中
	en_str_5 := "HELLO WORLD HUNAN CHANGSHA LUOBOZI"
	en_str_6 := "from changsha hunan china"
	fmt.Println("将en_str_5转换成小写：", strings.ToLower(en_str_5))
	fmt.Println("将en_str_6转换成大写：", strings.ToUpper(en_str_6))
	//写一个大小写转换的代码
	en_str_1 := "START|asdfljklFASDASDFasdfiaASDF123aszASDQWEQ|end"
	resp := ""
	for _, v := range en_str_1 {
		if v >= 97 && v <= 122 {
			resp += string(v - 32)
		} else if v >= 65 && v <= 90 {
			resp += string(v + 32)
		} else {
			resp += string(v)
		}
	}
	fmt.Println("en_str_1字母大小写转换后：", resp)

	//Trim(s1,s2 string类型) 将s1两侧的s2字符去除  常用方式：清洗首尾的空格、清洗首尾的特殊符号
	en_str_7 := "**这是md加粗语法**"
	fmt.Println("将en_str_7的前后的 * 号清洗掉：", strings.Trim(en_str_7, "*"))

	//HasPrefix(s1,s2 string类型) 判断s1是不是以s2开头 返回bool类型
	//HasSuffix(s1,s2 string类型) 判断s1是不是以s2结尾 返回bool类型
	en_str_8 := "IP：192.168.100.254"
	fmt.Println("en_str_8 是否以IP 开头：", strings.HasPrefix(en_str_8, "IP"))
	fmt.Println("en_str_8 是否以254 结尾：", strings.HasSuffix(en_str_8, "254"))

	var x bool
	fmt.Println("bool类型不赋值默认情况下为：", x)

	//GO的基本数据类型： 整形、浮点型、复数、字符串、字符型、布尔型
}
