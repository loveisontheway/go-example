package main

import "fmt"

// *** 函数返回值 ***
// "_"标识符，用来忽略函数的某个返回值
// Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。
// 返回值的名称应当具有一定的意义，可以作为文档使用。
// 没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。
// 直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
func main()  {
	var a, b int = 1, 2
	c := add(a, b)
	sum, avg := calc(a, b)
	fmt.Println(a, b, c, sum, avg)
	fmt.Println("--------------------")

	// Golang返回值不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略。
	// s := make([]int, 2)
	// s = testReturn()   // Error: multiple-value test() in single-value context

	x, _ := testReturn()
	println(x)
	fmt.Println("--------------------")
	// 多返回值可直接作为其他函数调用实参。
	println(add2(test2()))
	println(sum2(test2()))
	fmt.Println("--------------------")

	// 命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回。
	println(add3(1, 2))
}

func add(a, b int) (c int) {
	c = a + b
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}

func testReturn() (int, int) {
	return 1, 2
}

func test2() (int, int) {
	return 1, 2
}

func add2(x, y int) int {
	return x + y
}

func sum2(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}

func add3(x, y int) (z int) {
	z = x + y
	return
}

