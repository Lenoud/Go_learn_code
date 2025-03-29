// @FileName : 02文件操作.go
// @Time : 2025/3/19 下午2:04
// @Author : luobozi

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 此处开始编写代码
/*
文件操作
打开文件 -- os.Open os.OpenFile
os.Open
os.OpenFile 可以指定文件的打开方式，权限等等
*/

//func main() {
//	//openfile_read() //openfile方式去读取文件内容
//	//ioutil_read()
//	write_file()
//}

func openfile_read() {
	//OpenFile 指定文件的打开方式
	//O_RDONLY 只读
	//O_WRONLY 只写
	//o_RDWR 读写
	file, e := os.OpenFile("text.txt", os.O_RDONLY, 777) //默认是读取项目根目录下的txt文件
	if e != nil {
		fmt.Println("打开文件失败！")
	} else {
		//	data := make([]byte, 10)  //按照字节读取，每次10个字节
		//	for {
		//		len, _ := file.Read(data) //每次读取内容的长度，就是data的大小，适用于大文件的读取
		//		if len == 0 {             //如果长度为0表示读取完成
		//			break
		//		}
		//		fmt.Println("每次读取：", string(data))
		//	}
		//}
		//file.Close() //关闭打开文件，释放占用的内存资源

		//使用bufio读取  方式灵活  可以指定按行读取或者按分隔符读取 速度相对慢一点
		reader := bufio.NewReader(file)
		//按行读取
		for {
			line, _, e := reader.ReadLine()
			if e == io.EOF {
				break
			}
			fmt.Println("按行读：", string(line))
		}
		file.Seek(0, 0) //移动文件指针到最开始位置
	}
	file.Close()
}

func ioutil_read() {
	//使用ioutil读取 速度快   一次性读完 适合小文件读取  未来可能被淘汰
	file3, _ := ioutil.ReadFile("text.txt")
	fmt.Println(string(file3))
}

func write_file() {
	//写文件 io.WriteString  f.WriteString  bufio.NewWriter
	file, _ := os.OpenFile("text.txt", os.O_APPEND, 777)
	n, _ := io.WriteString(file, "\nthis is io.WriteString")
	fmt.Println("n是字节数：", n)
	file.WriteString("\n直接使用打开文件写入！")

}
