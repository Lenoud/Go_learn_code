// @FileName : 03继承和多态.go
// @Time : 2025/3/21 下午2:28
// @Author : luobozi

package main

import "fmt"

// 此处开始编写代码

/*
继承的实现: 结构体中包含另一个结构体
*/

type Person1 struct {
	name string
	age  int
}

type Person2 struct {
	name string
	age  int
}

type Teacher struct {
	Person1 //继承
	//Person2 //继承多个
	subject string
	name    int //重写
}

func (p *Person1) Hello() {
	fmt.Println("i am Person1....")
}

//func main() {
//	//t1 := Teacher{}
//	//
//	//t1.Hello() //看上去就像Teacher实例化后继承了Persion1父类的方法
//	////语法糖 实际上是下面这种
//	//t1.Person1.Hello()
//	//t1.name = 100               //看上去就像，子类有了就不会去找父类的属性了
//	//t1.Person1.name = "luobozi" //但是可以通过这种方式来操作父类的属性
//	//fmt.Printf("%+v\n", t1)
//
//	//继承多个 属性字段如果发生冲突 必须分开 (和python的继承不同，不会自动判断顺序)
//	//t1.Person1.age = 18
//	//t1.Person2.age = 21
//	//fmt.Printf("%+v\n", t1)
//
//	//实例化
//	t2 := Teacher{
//		Person1: Person1{"zhangsan", 12},
//		subject: "math",
//	}
//	fmt.Printf("%+v\n", t2)
//
//	//先实例化父类，将对象赋值给子类
//	p1 := Person1{"mikodo", 19}
//	t3 := Teacher{Person1: p1, subject: "golang"}
//	fmt.Printf("%+v\n", t3)
//
//}

//多态
/*
		go语言实现多态： interface类型去实现
		go语言多态：也和python一样崇尚鸭子类型，不关心具体类型，只关心有没有这个行为

	定义接口：
	type 接口名 interface{
	方法1() []
	方法2() []
	.....
	}
*/
type MoneyPay interface {
	pay()
}

// 空接口 : 可以代表任何数据类型
type kong interface {
}
type ZFB struct {
	name string
}
type WX struct {
	name string
}

func (z *ZFB) pay() {
	fmt.Println("this is zfb pay")
}

func (w *WX) pay() {
	fmt.Println("this is wx pay")
}

func main() {
	z1 := &ZFB{"支付宝"}
	w1 := &WX{"微信"}

	var j1, j2 MoneyPay

	j1 = z1
	j2 = w1
	fmt.Println(j1, j2)

	FinPay(z1)
	FinPay(w1)

	//空接口 代表任何数据类型
	m1 := map[string]kong{
		"name": "luobozi",
		"age":  18,
	}
	fmt.Println(m1)
	//类型转换
	result := m1["age"].(int) + 10
	fmt.Println(result)
}

//函数接口

func FinPay(p MoneyPay) {
	p.pay()
}
