// @FileName : 02切片.go
// @Time : 2025/3/13 下午4:49
// @Author : luobozi

package main

import "fmt"

// 此处开始编写代码
//go数组的长度固定不可改变，在特定的场景中这样的集合就有局限性
//go提供了一种更加灵活的内置类型 slice -- 切片  （"动态数组"）
//与数组相比，切片长度不固定，可以追加删除元素

// var name []type
// name := []type{1,2,3}

//使用make函数定义切片  make作用 -->分配内存空间
//var name []type =make([]type,len,cap)
//name:=make([]type,len,cap)
//使用len()查看切片长度
//使用cap()查看你容量大小

//切片不存储具体的值，切片存储的是  指向的底层数组的地址，和长度、容量大小

func main() {
	sli_1()
}

func sli_1() {
	var lst1 []string
	lst1 = append(lst1, "abc") //往切片中增加元素
	lst1 = append(lst1, "def")
	fmt.Println(lst1)

	lst2 := []string{"asdfas", "sc", "haha"}
	fmt.Println(lst2)

	lst3 := make([]int, 5)
	fmt.Println(lst3)

	lst4 := make([]int, 7, 10)
	fmt.Println(lst4)
	fmt.Println("长度为：", len(lst4))
	fmt.Println("容量为：", cap(lst4))

	//切片的截取
	//从一个切片创建另一个切片
	//slice[start:end:cap]  cap需要超过或者等于end
	lst5 := []int{1, 2, 3, 4, 5, 6, 7}
	lst6 := lst5[1:3:5]
	fmt.Printf("lst5地址：%p\n", lst5)
	fmt.Printf("lst6地址：%p\n", lst6)
	lst6[1] = 200 //修改lst6的值也会改动lst5，因为指针指向同一个地址
	fmt.Println(lst5)
	fmt.Println(lst6)

	//切片拷贝
	lst7 := []int{4, 5, 6, 7, 8}

	//lst8 := lst7 //这样子赋值，会指向同一个底层数组
	lst8 := make([]int, 5)
	copy(lst8, lst7)
	lst7[0] = 100
	fmt.Print(lst7) //可以看到两个切片互不影响
	fmt.Println(lst8)

	//删除  go里面没有直接删除的操作，只能使用巧方式 将对应的下标 数据删除
	//比如我们要删除lst8 中的下标1  的值 ：
	lst8 = append(lst8[:1], lst8[2:]...) // ... 解包操作等效于  append(lst8[:1], 6,7,8)  [6 7 8] -> 6,7,8
	fmt.Println(lst8)

	//nil 和 空切片
	//在go中 nil 和空切片 是不同的概念，空切片已被分配了内存地址，但是nil 是指针指向了nil还没有分配内存地址
	var lst9 []int
	if lst9 == nil {
		fmt.Println("lst9为nil")
	}

	//s1:=make([]int,0)
	s1 := []int{}  //空切片 --> 已经初始化了  空切片指针指向了一个内存地址 但是没有分配空间
	if s1 != nil { //所以要想判断是否为空指针 那么需要使用 len == 0的方式来判断
		fmt.Println("s1是空指针,但是不是nil")
	}
}
