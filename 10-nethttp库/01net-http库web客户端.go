// @FileName : 01net-http库web客户端.go
// @Time : 2025/3/22 下午7:58
// @Author : luobozi

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 此处开始编写代码
/*
net/http 库 go语言的标准库
作为http客户端，它可以发起http请求 可以请求比尔的网址接口 类似python中的requests
作为http服务端，还可以启动一个web服务，提供页面展示
*/
func main() {
	//客户端 发起请求
	request_client("https://www.baidu.com")
	download_pic("https://www.baidu.com/img/bd_logo1.png")
}

func request_client(url string) {
	//创建一个客户端  使用客户端连接
	client := http.Client{}
	//创建一个新的请求 规定使用什么方法去请求那个url 传输哪些数据过去
	// http  --- post delete put get
	r, _ := http.NewRequest("GET", url, nil)
	//客户端发起请求
	response, e := client.Do(r)
	if e != nil {
		fmt.Println("error!：", e)
		return
	}
	fmt.Println("请求的状态码：", response.StatusCode)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取内容失败！", err)
		return
	}
	fmt.Println(string(body))
}

// 下载图片
func download_pic(url string) {
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		fmt.Println("下载图片错误：", err)
		return
	}
	//下载图片写入的文件
	file, err := os.Create("image01.png")
	if err != nil {
		fmt.Println("写入图片错误：", err)
		return
	}
	//将respones.Body 图片数据写入本地文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("图片保存错误：", err)
		return
	}
}
