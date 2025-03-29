// @FileName : 02net-http库-服务端.go
// @Time : 2025/3/22 下午8:22
// @Author : luobozi

package main

import (
	"fmt"
	"net/http"
)

// 此处开始编写代码
/*
服务端实现

编写web网站
路由： url 视图函数

*/
func main() {
	//路由绑定
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/register", register)
	err := http.ListenAndServe("localhost:9000", nil)
	if err != nil {
		fmt.Println("发生错误！", err)
	}
	
}

// 视图函数 处理和响应客户端请求的函数
// r 客户端的请求信息
// w 服务端的响应数据
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("这是首页"))

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("这是home"))

}
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>这是hello</h1>"))
}
func register(w http.ResponseWriter, r *http.Request) {
	//获取url传参
	parm := r.URL.Query()
	name := parm.Get("name")
	passwd := parm.Get("passwd")
	fmt.Println("this is name:", name)
	fmt.Println("this is passwd:", passwd)
	body := []byte(name + "==" + passwd)
	w.Write(body)
}
