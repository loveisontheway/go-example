package main

import "fmt"

// *** Channel ***
// 单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。
//
// 虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。
// 为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。
//
// Go语言的并发模型是CSP（Communicating Sequential Processes），
// 提倡通过通信共享内存而不是通过共享内存而实现通信。
//
// 如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。
// channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。
//
// Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，
// 总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
// 每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
func main() {

	// channel类型
	// channel是一种类型，一种引用类型。声明通道类型的格式如下
	// var 变量 chan 元素类型
	var ch1 chan int   // 声明一个传递整型的通道
	var ch2 chan bool  // 声明一个传递布尔型的通道
	var ch3 chan []int // 声明一个传递int切片的通道
	fmt.Println(ch1, ch2, ch3)
	fmt.Println("---------------------")

	// 创建channel
	// 通道是引用类型，通道类型的空值是nil。
	var ch chan int
	fmt.Println(ch) // <nil>
	fmt.Println("---------------------")
	// 声明的通道后需要使用make函数初始化之后才能使用。
	// 创建channel的格式如下：
	// make(chan 元素类型, [缓冲大小])
	// channel的缓冲大小是可选的。
	ch4 := make(chan int)
	ch5 := make(chan bool)
	ch6 := make(chan []int)
	fmt.Println(ch4, ch5, ch6)
	fmt.Println("---------------------")

	// channel操作
	// 通道有发送（send）、接收(receive）和关闭（close）三种操作。
	// 发送和接收都使用<-符号。
	// 现在我们先使用以下语句定义一个通道：
	ch7 := make(chan int)

	// 发送
	// 将一个值发送到通道中。
	ch7 <- 10 // 把10发送到ch7中

	// 接收
	// 从一个通道中接收值。
	x := <-ch7 // 从ch中接收值并赋值给变量x
	<-ch7      // 从ch中接收值，忽略结果

	// 关闭
	// 我们通过调用内置的close函数来关闭通道。
	close(ch7)
	fmt.Println(x)
	/*
		关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。
		通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，
		在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
	关闭后的通道有以下特点：
	    1.对一个关闭的通道再发送值就会导致panic。
	    2.对一个关闭的通道进行接收会一直获取值直到通道为空。
	    3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
	    4.关闭一个已经关闭的通道会导致panic。
	 */
	fmt.Println("---------------------")
}
