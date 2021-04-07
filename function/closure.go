package main

import "fmt"

// *** 闭包 ***
func main()  {
	c := a()
	c()
	c()
	c()

	a() //不会输出i
}

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}
