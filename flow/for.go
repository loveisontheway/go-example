package main

import "fmt"

// *** 循环语句for ***
// Golang for支持三种循环方式，包括类似 while 的语法。
// for循环是一个循环控制结构，可以执行指定次数的循环。
// Go语言的For循环有3中形式，只有其中的一种使用分号。
/*
	for init; condition; post { }
    for condition { }
    for { }
    init： 一般为赋值表达式，给控制变量赋初值；
    condition： 关系表达式或逻辑表达式，循环控制条件；
    post： 一般为赋值表达式，给控制变量增量或减量。
    for语句执行过程如下：
    ①先对表达式 init 赋初值；
    ②判别赋值表达式 init 是否满足给定 condition 条件，
     若其值为真，满足循环条件，则执行循环体内语句，
     然后执行 post，进入第二次循环，
     再判别 condition；否则判断 condition 的值为假，
     不满足条件，就终止for循环，执行循环体外语句。
 */
func main()  {
	/*
	s := "abc"

	for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
		println(s[i])
	}

	n := len(s)
	for n > 0 {              // 替代 while (n > 0) {}
		println(s[n])        // 替代 for (; n > 0;) {}
		n--
	}

	for {                    // 替代 while (true) {}
		println(s)            // 替代 for (;;) {}
	}
	*/

	s := "abcd"

	for i, n := 0, length(s); i < n; i++ {     // 避免多次调用 length 函数。
		println(i, s[i])
	}

	fmt.Println("----------------------")
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
	fmt.Println("-----------------")

	// 循环嵌套
	// 在 for 循环中嵌套一个或多个 for 循环
	/*
	for [condition |  ( init; condition; increment ) | Range]
	{
	   for [condition |  ( init; condition; increment ) | Range]
	   {
	      statement(s)
	   }
	   statement(s)
	}
	 */
	/* 定义局部变量 */
	var i, j int

	for i=2; i < 100; i++ {
		for j=2; j <= (i/j); j++ {
			if(i%j==0) {
				break // 如果发现因子，则不是素数
			}
		}
		if(j > (i/j)) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	fmt.Println("-----------------")

	// 无限循环
	for true  {
		fmt.Printf("这是无限循环。\n")
	}
}

func length(s string) int {
	println("call length.")
	return len(s)
}
