package main

import (
	"fmt"
	"math"
)

// *** 匿名函数 ***
// 匿名函数是指不需要定义函数名的一种函数实现方式。1958年LISP首先采用匿名函数。
// 在Go里面，函数可以像普通变量一样被传递或使用，Go语言支持随时在代码里定义匿名函数。
// 匿名函数由一个不带函数名的函数声明和函数体组成。
// 匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
func main()  {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))
	fmt.Println("------------------")
	// --- function variable ---
	fn := func() { println("Hello, World!") }
	fn()

	// --- function collection ---
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	println(fns[0](100))

	// --- function as field ---
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	println(d.fn())

	// --- channel of function ---
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	println((<-fc)())
}
