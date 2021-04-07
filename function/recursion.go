package main

import "fmt"

// *** 递归函数 ***
// 递归，就是在运行的过程中调用自己。
// 一个函数调用自己，就叫做递归函数。
// 构成递归需具备的条件：
// 1.子问题须与原始问题为同样的事，且更为简单。
// 2.不能无限制地调用本身，须有个出口，化简为非递归状况处理。
func main()  {
	// 数字阶乘
	// 一个正整数的阶乘（factorial）是所有小于及等于该数的正整数的积，
	// 并且0的阶乘为1。自然数n的阶乘写作n!。1808年，基斯顿·卡曼引进这个表示法。
	var i int = 7
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
	fmt.Println("---------------------")

	// 斐波那契数列(Fibonacci)
	var j int
	for j = 0; j < 10; j++ {
		fmt.Printf("%d\n", fibonaci(j))
	}
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}
