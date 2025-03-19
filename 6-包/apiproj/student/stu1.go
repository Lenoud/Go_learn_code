// @FileName : stu1.go
// @Time : 2025/3/18 下午5:17
// @Author : luobozi

package mystudent //包名和目录名可以不一样，目录下的文件包名只能有一个，要统一
import "fmt"

// 此处开始编写代码
var stuname string = "luobozi" //同一个包下面不同文件中的函数、变量可以随便调用
//staname 识别成私有成员，只能在包内使用。别的包不能调用
//标识符在包中，如果首字母小写，识别成私有成员，只能在包中使用
//			 如果首字母大写。识别成公有成员，可以在其它包中使用

var Stuage int = 18

func init() {
	fmt.Println("this is stu1.go init...")
	//这个函数在导入包的时候自动执行
	//一般用于做一些初始化工作
}

func Test() {
	fmt.Println(Stuage)
}
