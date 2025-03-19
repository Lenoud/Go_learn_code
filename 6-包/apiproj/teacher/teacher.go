// @FileName : teacher.go
// @Time : 2025/3/18 下午5:34
// @Author : luobozi

package teacher

import "fmt"

func init() {
	fmt.Println("这是包中的特殊函数，init")
	//这个函数在导入包的时候自动执行
	//一般用于做一些初始化工作
}

func init() {
	fmt.Println("init ... 2")
	//这个函数在导入包的时候自动执行
	//一般用于做一些初始化工作
}

// 此处开始编写代码
func TeacherInfo() {
	fmt.Println("this is teacher info")

}
