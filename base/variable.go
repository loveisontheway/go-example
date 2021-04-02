package main

import "fmt"

// 批量声明变量
var (
	a string
	b int
	c bool
	d float32
)

// 全局变量
var m = 100

// *** 变量 ***
// Go语言的变量声明后必须使用
// Go语言的变量声明格式为：var 变量名 变量类型
func main() {

	// 变量声明以关键字var开头，变量类型放在变量的后面，行尾无需分号。 举个例子：
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)
	fmt.Println(a, b, c, d)

	// 变量初始化
	var email string = "xx@gmail.com"
	var sex int = 1
	fmt.Println(email, sex)

	// 一次初始化多个变量
	var address, num = "xxx.china", 90
	fmt.Println(address, num)

	// 类型推导
	// 将变量的类型省略，编译器会根据等号右边的值来推导变量的类型完成初始化
	var userName = "muxi"
	var grade = 1
	fmt.Println(userName, grade)

	// 短变量声明
	n := 10
	m := 200	// 此处声明局部变量m
	fmt.Println(m,n)

	// 匿名变量
	// 在使用多重赋值时，如果想要忽略某个值，
	// 可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示
	x,_ := foo()
	_,y := foo()
	fmt.Println("x=",x,"\ty=",y)
}

func foo() (int, string) {
	return 10, "Ximi"
}
