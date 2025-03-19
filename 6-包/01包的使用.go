// @FileName : 1包的使用.go
// @Time : 2025/3/18 下午3:11
// @Author : luobozi

package main

// 此处开始编写代码

// 自定义包
// 内置包
// 第三方包
//1.go语言模块文件的顶部，声明这是哪个包
//2.每个程序 必须包含一个main包，main包中要有main函数，程序入口就是main函数
//3.go语言一个目录下放一个包，一个包可以有不同的文件，目录的名字可以跟包名不一样
//4.导入了包，就一定要使用这个包
//5.包中的标识符首字母大写表示公有成员（可以给外部包调用），首字母没有大写表示私有成员只能在包中调用

/*
GO命令
(base) PS C:\Users\Administrator> go env
set GO111MODULE=on  //go项目包管理方式
set GOPATH=D:\Environment\GO\gopath   //go的包管理路径
set GOPROXY=https://goproxy.cn,direct   //下载第三方包的代理
set GOROOT=D:\Environment\GO  //go的安装路径

GO111MODULE  go项目包管理方式
GOPATH  go的包管理路径
GOPROXY 下载第三方包的代理
GOROOT  go的安装路径

GOROOT/src 内置包一般放在这个下面

# 修改go环境变量
#修改成国内代理网址
go env -w  GOPROXY=https://goproxy.cn,direct   //七牛云代理网址

GO包管理方式
1.gopath
早期的go项目必须放在$GOPATH/src目录下，下载的三方包也放在$GOPATH/pkg目录下
这种管理会造成多个项目混用，管理起来不方便，也不好控制版本
2.go vender
项目还是放在goPath目录下面，引入了一个vendor可以进行不同项目的版本管理
3.go mod
项目工程目录可以放在goPath路径之外
要求项目中必须要有go.mod文件，go.mod文件记录当前项目需要的第三方软件以及它的版本
所有的第三方依赖包都放在$gopath/pkg/mod 目录下
go env -w GO111MODULE=on 开启gomod管理方式

go mod 命令
go mod init xxx 使用go mod 方式管理项目，初始化项目xxx   xxx的名字为项目包名

go mod tidy  go自动检测缺少哪些包，自动安装
go get github.com/gin-gonic/gin 手动安装包
go mod download

go -- gin框架 做web 开发

init 函数 包初始化函数，导入包的时候自动执行init函数
init函数用于包执行之前做初始化工作
每个包可以拥有多个init函数顺序执行
不同包的init函数按照包的导入依赖关系决定init的执行顺序
init函数不能被其它函数调用

开启go mod 后包的查找顺序：
现在当前目录下找
然后去找GOPATH/pkg/mod 目录查找相应的版本进行导入
如果还没找到，再去GOPATH/src下查找
最后去GOROOT/src目录查找
*/
func main() {

}
