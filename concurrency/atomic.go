package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// *** 原子操作(atomic包) ***
// 原子操作
// 代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。
// 针对基本数据类型我们还可以使用原子操作来保证并发安全，
// 因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好。
// Go语言中原子操作由内置的标准库sync/atomic提供。
func main()  {

	// 我们填写一个示例来比较下互斥锁和原子操作的性能。
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg6.Add(1)
		// go add6()       // 普通版add函数 不是并发安全的
		// go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg6.Wait()
	end := time.Now()
	fmt.Println(x6)
	fmt.Println(end.Sub(start))
}

var x6 int64
var l6 sync.Mutex
var wg6 sync.WaitGroup

// 普通版加函数
func add6() {
	// x6 = x6 + 1
	x6++ // 等价于上面的操作
	wg6.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l6.Lock()
	x6++
	l6.Unlock()
	wg6.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x6, 1)
	wg6.Done()
}
