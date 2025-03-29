// @FileName : 04sync.go
// @Time : 2025/3/23 下午8:08
// @Author : luobozi

package main

import (
	"fmt"
	"sync"
	"time"
)

// 此处开始编写代码
/*
go语言sync包 提供一些用于并发操作的同步原语
常见的锁机制
	互斥锁Mutex
	条件变量Cond
	WaitGroup  阻塞
*/

var wg sync.WaitGroup
var num = 10
var mutex sync.Mutex //互斥锁
func main() {
	wg.Add(3)
	go test(&wg)
	go test(&wg)
	go test(&wg)
	wg.Wait()
	fmt.Println("执行结束")
}

func test(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock() //获取锁
	time.Sleep(1 * time.Second)
	fmt.Println("this is test...")
	num += 1
	fmt.Println("num:", num)
	mutex.Unlock() //释放
}
