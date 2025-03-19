// @FileName : 03map.go
// @Time : 2025/3/13 下午9:08
// @Author : luobozi

package main

import (
	"encoding/json"
	"fmt"
)

// 此处开始编写代码
func main() {
	map_1()
}

func map_1() {
	//key value 键值存储的数据结构
	// map 引用类型  保存底层hash桶  key 天生去重 可hash对象 值类型
	//var dict map[string]string
	//引用类型： 不保存具体值，只保存指向底层存储的指针
	//值类型：保存具体的值
	//go语言中引用类型：切片、map、channel这些都是
	//go语言中值类型：整型、浮点型、复数、字符串、数组

	//创建和新增值
	dict_1 := make(map[string]string, 100)
	dict_1["china"] = "中国"
	dict_1["changsha"] = "长沙"
	fmt.Println(dict_1["china"])

	if v, ok := dict_1["china"]; ok {
		fmt.Println("key存在", v, ok)
	} else {
		fmt.Println("key不存在", ok)
	}

	result := dict_1["changsha"]
	fmt.Println(result)

	//声明加赋值
	dict2 := map[string]string{
		"china":    "中国",
		"changsha": "长沙",
		"hunan":    "湖南",
	}
	fmt.Println(dict2)

	//遍历
	//用一个变量接收，只会遍历key
	for k := range dict2 {
		fmt.Println(k, "is", dict2[k])
	}
	//使用两个变量去接收
	for k, v := range dict2 {
		fmt.Println(k, "--->", v)
	}

	//删除
	delete(dict2, "china")
	fmt.Println(dict2)

	//map包map套娃
	UserInfo := map[string]map[string]string{
		"root":    {"password": "123456"},
		"luobozi": {"password": "2024022236"},
	}
	fmt.Println(UserInfo)
	fmt.Println(UserInfo["luobozi"]["password"])

	//空接口  -- 可以代表任何类型对象
	emptyMap := map[string]interface{}{
		"root": "root123",
		"age":  18,
	}

	fmt.Println("空接口age值：", emptyMap["age"])
	fmt.Printf("空接口类型：%T\n", emptyMap["age"])
	fmt.Println("空接口类型转换才能进行操作：", emptyMap["age"].(int)+2) //空接口类型进行转换才能和其它数据类型进行转换

	fmt.Println("map，json之间互转操作")
	//map和json格式的转换
	userStr, e := json.Marshal(UserInfo)
	if e != nil {
		fmt.Println("转换出错")
	} else {
		fmt.Println("转换成功")
		fmt.Printf("转换后的数据类型：%T\n", userStr)
		fmt.Println(userStr)
		fmt.Println(string(userStr))
	}

	//json -->map
	data_str := `{"a":1,"b":2}`
	fmt.Println(data_str)
	data_map := map[string]int{}
	err := json.Unmarshal([]byte(data_str), &data_map)
	if err != nil {
		fmt.Println("转换出错", err)
	} else {
		fmt.Println(data_map)
	}
}
