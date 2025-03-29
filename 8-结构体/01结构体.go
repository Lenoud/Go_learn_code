// @FileName : 01结构体.go
// @Time : 2025/3/20 下午4:59
// @Author : luobozi

package main

import "fmt"

// 此处开始编写代码
/*
go语言不是面向对象的语言
结构体  类似于面向对象
定义结构体：
type 自定义类型 struct {
  field_name type
  field_name type
  ......
}
为结构体绑定方法：
func (s RecevierType) Funcname(parameters) (results){
	....
}

//支持匿名结构体  不需要type关键字定义，定义结构体的时候就需要实例化对象

变量名:= struct{
	结构题成员属性
}{
	初始化成员属性
}

//匿名字段属性
type Person struct{
	string
	int
}

//空结构体：  占位，声明出来不占用内存
一般用在
1.实现方法接收者
2.集合类型  在python中就是只有key的字典，key 必须是不可重复、不可变、可哈希对象
3.实现空通道，做协程控制
*/

// 定义一个结构体
type Student struct {
	id     int
	name   string
	age    int
	school string
}

// 空结构体
type Temp struct{}

func (t *Temp) Call() {
	fmt.Println("call for temp")
}

// 空结构体 实现集合 只有key的map
type Set map[string]struct{} //使用map为基本，创建set类型
// 往集合里面添加元素
func (s Set) Append(k string) {
	s[k] = struct{}{} //空结构体是不占用内存的，所以就用这种方式实现集合，本质上也是一个map
}

func main() {
	//实例化（指针类型）
	pan := &Student{ //结构体名字前面加 & 就是实例化指针类型
		id:     1,
		name:   "luobozi",
		age:    18,
		school: "znlsw",
	}
	fmt.Println((*pan).name)
	// 语法糖 简化开发
	fmt.Println(pan.name)
	fmt.Println()
	pan.age = 20
	fmt.Println(pan)
	fmt.Printf("%+v\n", pan) //展示结构体完整信息

	//使用new 进行实例化（指针类型）
	//new是一个分配内存空间的函数，但是它返回的都是指针类型，常常用于对结构体的实例化工作
	//new和make的区别，他们都是分配内存空间的函数，new返回的是分配内存空间的地址，而make返回的是保存的具体的值
	//new是用来做结构体的初始化，make是用来做切片mapchnale的初始化
	luo := new(Student)
	luo.id = 2
	luo.name = "luobozi"
	fmt.Printf("%+v\n", luo)

	//实例化对象（值类型）
	bo := Student{
		id:   3,
		name: "luobozi",
		age:  18,
	}
	fmt.Printf("%+v\n", bo)

	//传递赋值
	//指针类型
	luo2 := luo
	luo.id = 999
	fmt.Println(luo2.id, luo.id)

	//值类型 ，会拷贝一份数据，修改不影响其它变量
	bo2 := bo
	bo2.name = "张三"
	fmt.Println(bo2.name, bo.name)

	test(bo)
	fmt.Printf("%+v\n", bo)

	test1(luo)
	fmt.Printf("%+v\n", luo)
	fmt.Println("")

	//调用方法
	bo.ChangeName("milamn")
	fmt.Println(bo)
	luo.ChangeName("luji")
	fmt.Println(luo)

	bo.ChangeName2("wesiman")
	fmt.Println(bo)
	luo.ChangeName2("mikod")
	fmt.Println(luo)
	fmt.Println()
	//定义匿名结构体
	abc := struct {
		abcid   int
		abcaddr string
	}{
		abcid:   10,
		abcaddr: "长沙",
	}
	fmt.Println(abc.abcaddr)
	fmt.Println()

	//匿名字段
	qwe := Person{
		"niming", 2,
	}
	fmt.Println(qwe)
}

// 传递值类型
func test(s Student) {
	fmt.Println("test change age...传递值")
	s.age = 100
}

// 传递指针
func test1(s *Student) {
	fmt.Println("test1 change age...传递指针")
	s.age = 200
}

// 为Student 结构体添加方法
func (s Student) ChangeName(name string) {
	fmt.Println("调用方法ChangeName")
	s.name = name
}

func (s *Student) ChangeName2(name string) {
	fmt.Println("调用方法ChangeName2")
	s.name = name
}

// 匿名字段属性
type Person struct {
	string
	int
}

//内存对齐
