package main

import (
	"fmt"
	"sync"
	"time"
)

// *** Goroutine ***
// Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，
// 但 goroutine是由Go的运行时（runtime）调度和管理的。
// Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。
// Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。
//
// 在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine，
// 当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，
// 开启一个goroutine去执行这个函数就可以了，就是这么简单粗暴。

// 使用goroutine
// Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，
// 就可以为一个函数创建一个goroutine。
//
// 一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。

// 启动单个goroutine
// 启动goroutine的方式非常简单，只需要在调用的函数（普通函数和匿名函数）前面加上一个go关键字。
func main()  {
	// 执行结果只打印了main goroutine done!，并没有打印Hello Goroutine!。为什么呢？
	// 在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。
	// 当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束，
	// main函数所在的goroutine就像是权利的游戏中的夜王，
	// 其他的goroutine都是异鬼，夜王一死它转化的那些异鬼也就全部GG了。
	go hello()
	fmt.Println("main goroutine done!")
	// 所以我们要想办法让main函数等一等hello函数，最简单粗暴的方式就是time.Sleep了。
	time.Sleep(time.Second)
	fmt.Println("---------------------")
	/*
	执行上面的代码你会发现，这一次先打印main goroutine done!，然后紧接着打印Hello Goroutine!。
	首先为什么会先打印main goroutine done!是因为我们在创建新的goroutine的时候需要花费一些时间，
	而此时main函数所在的goroutine是继续执行的。
	*/

	// 启动多个goroutine
	// 在Go语言中实现并发就是这样简单，我们还可以启动多个goroutine。
	// 让我们再来一个例子： （这里使用了sync.WaitGroup来实现goroutine的同步）
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello2(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
	fmt.Println("---------------------")
	/*
	多次执行上面的代码，会发现每次打印的数字的顺序都不一致。
	这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。
	 */

	// 注意
	// 如果主协程退出了，其他任务还执行吗（运行下面的代码测试一下吧）
	// 合起来写
	go func() {
		j := 0
		for {
			j++
			fmt.Printf("new goroutine: i = %d\n", j)
			time.Sleep(time.Second)
		}
	}()
	j := 0
	for {
		j++
		fmt.Printf("main goroutine: i = %d\n", j)
		time.Sleep(time.Second)
		if j == 2 {
			break
		}
	}
}

func hello() {
	fmt.Println("Hello Goroutine!")
}

var wg sync.WaitGroup
func hello2(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
