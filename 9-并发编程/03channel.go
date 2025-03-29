// @FileName : 03channel.go
// @Time : 2025/3/23 下午2:17
// @Author : luobozi

package main

//
//import (
//	"fmt"
//	"time"
//)
//
//// 此处开始编写代码
///*
//CSP: 不要以共享内存的方式来通信 相反 要通过通信来共享内存
//channel 本身就是一个队列，先进先出
//channel 是引用类型 可以使用make初始化 make: map、slice、channel
//channel 写满的时候不能写，取空的时候不能取，很容易发生堵塞，很容易产生死锁
//每个channel 只能放同一种数据类型
//使用channel 收发操作都是在不同的goroutine
//
//channel
//	无缓冲通道
//	有缓冲通道
//*/
//
//func main() {
//	//无缓冲通道
//	ch1 := make(chan int)
//	go senddata(ch1)
//	//ch1 <- 2 报错，因为只能不同协程的操作
//	time.Sleep(3 * time.Second)
//	data := <-ch1
//	fmt.Println("拿到通道数据：", data)
//	data2 := <-ch1
//	fmt.Println("拿到通道数据：", data2)
//	data3 := <-ch1
//	fmt.Println("第三次拿到通道数据：", data3)
//	//time.Sleep(3 * time.Second)
//}
//func senddata(ch chan int) {
//	ch <- 1
//	fmt.Println("往通道里存放数据1")
//	ch <- 2
//	fmt.Println("往通道里存放数据2")
//}
//
//// 利用无缓冲通道，做协程控制
//// 主协程退出，子协程也会跟着退出。利用通道去控制主协程等待子协程执行完毕再退出
//type none struct{}
//
//func senddata(ch chan none) {
//	fmt.Println("this is senddata")
//	time.Sleep(2 * time.Second)
//	ch <- none{} //空结构体不占用任何空间
//}
//func main() {
//	fmt.Println("start....")
//	ch1 := make(chan none)
//	go senddata(ch1)
//	<-ch1
//	fmt.Println("end....")
//}
//
//// 通道的关闭
//func sendata(ch chan string) {
//	for i := 0; i < 3; i++ {
//		ch <- fmt.Sprintf("发送数据%d", i)
//	}
//	defer fmt.Println("发送数据完毕")
//	defer close(ch) //数据发送完毕关闭通道,拿去数据方不会因为没有数据而产生死锁
//}
//func main() {
//	ch1 := make(chan string)
//	go sendata(ch1)
//	//消费方判断通道是否关闭，来停止消费
//	for {
//		//方式1
//		data := <-ch1
//		//如果通道关闭，读取的就是数据类型的默认值 string "" int 0
//		if data == "" {
//			break
//		}
//		fmt.Println("从通道中获取到的数据：", data)
//	}
//
//	//方式2
//	for {
//		data, ok := <-ch1
//		//如果ok为false 表示通道关闭
//		if !ok {
//			break
//		}
//		fmt.Println("2从通道中获取到的数据：", data)
//	}
//
//	//方式3 range string slice map channel
//	//range 会自动判断通道是否关闭
//	for value := range ch1 {
//		fmt.Println("3从通道中获取到的数据：", value)
//	}
//}
//
//// 有缓冲通道
//func main() {
//	ch1 := make(chan string, 6)
//	go senddata(ch1)
//	time.Sleep(time.Second * 2)
//	for v := range ch1 {
//		fmt.Println("读取通道数据", v)
//	}
//
//}
//func senddata(ch chan string) {
//	for i := 0; i <= 10; i++ {
//		ch <- fmt.Sprintf("放入数据%d", i)
//		fmt.Printf("往通道放入数据%d\n", i)
//	}
//	defer close(ch)
//
//}
//
//// select 监听通道
//// select 可以同时监听多个通道操作，实现非阻塞通信和多路复用，实现高效的并发处理
//// select语句监听多个通道操作，并执行第一个准备好的操作。如果多个同时准备好，就会随机挑选一个执行
//// 如果没有通道准备好，select语句一直阻塞，直到至少有一个通道操作准备好
//// default 没有任何消息发生的时候执行，通过defaultl灵活处理非阻塞通道操作
//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	ch3 := make(chan int)
//
//	go func() {
//		time.Sleep(10 * time.Second)
//		ch1 <- 1
//	}()
//
//	go func() {
//		time.Sleep(2 * time.Second)
//		ch2 <- 2
//	}()
//
//	go func() {
//		time.Sleep(3 * time.Second)
//		ch3 <- 3
//	}()
//
//	for i := 0; i < 3; i++ {
//		select {
//		case msg := <-ch1:
//			fmt.Println("接收到数据from ch1", msg)
//		case msg := <-ch2:
//			fmt.Println("接收到数据from ch2", msg)
//		case msg := <-ch3:
//			fmt.Println("接收到数据from ch3", msg)
//		case <-time.After(1 * time.Second):
//			fmt.Println("设置超时时间！timeout")
//		default:
//			fmt.Println("没有消息准备好")
//		}
//
//	}
//}
