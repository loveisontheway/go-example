package main

import "fmt"

// *** 条件语句if ***
// • 可省略条件表达式括号。
// • 持初始化语句，可定义代码块局部变量。
// • 代码块左 括号必须在条件表达式尾部。
/*
	if 布尔表达式 {
		在布尔表达式为 true 时执行
	}
*/
// *不支持三元操作符(三目运算符) "a > b ? a : b"。
func main()  {
	x := 0

	// if x > 10        // Error: missing condition in if statement
	// {
	// }

	if n := "abc"; x > 0 {     // 初始化语句未必就是定义变量， 如 println("init") 也是可以的。
		println(n[2])
	} else if x < 0 {
		println(n[1])
	} else {
		println(n[0])
	}

	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n" )
		}
	}
	fmt.Printf("a 值为 : %d\n", a )
	fmt.Printf("b 值为 : %d\n", b )
}
