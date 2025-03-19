// @FileName : main.go.go
// @Time : 2025/3/18 下午5:16
// @Author : luobozi

package main

import (
	mystudent "apiprojct/student" //如果包名和目录名不一样就要这样声明
	//. "apiprojct/teacher"         //导入teacher包中的所有内容
	_ "apiprojct/teacher" //匿名导入，导入之后可以不适用，只需要执行包中的init函数
	//t "apiprojct/teacher"         //为导入的teacher包取别名
	"fmt"
)

// 此处开始编写代码
func main() {
	fmt.Println("这是main包main函数")
	mystudent.StuInfo()

	//私有和共有变量
	//fmt.Println(mystudent.stuname) //不能访问其它包的私有成员
	fmt.Println(mystudent.Stuage) //访问包中的公有成员
	mystudent.Stuage = 200        //修改公有成员
	mystudent.Test()              //查看修改后的公有成员
	//别名的使用
	//t.TeacherInfo()

	//导入teacher下的所有包,直接接函数就可以调用
	//TeacherInfo()
}
