package main

import "fmt"

// *** 方法定义 ***

// Golang 方法总是绑定对象实例，并隐式将实例作为第一实参 (receiver)。
/*
	• 只能为当前包内命名类型定义方法。
	• 参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名。
	• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。
	• 不支持方法重载，receiver 只是参数签名的组成部分。
	• 可用实例 value 或 pointer 调用全部方法，编译器自动转换。
 */
// 一个方法就是一个包含了接受者的函数，
// 接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
// 所有给定类型的方法属于该类型的方法集。
/*
	func (recevier type) methodName(参数列表)(返回值列表){}

    参数和返回值可以省略
 */
func main()  {
	// 值类型调用方法
	u1 := User{"golang", "golang@golang.com"}
	u1.Notify()
	// 指针类型调用方法
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()
}

type Test struct{}

// 无参数、无返回值
func (t Test) method0() {

}

// 单参数、无返回值
func (t Test) method1(i int) {

}

// 多参数、无返回值
func (t Test) method2(x, y int) {

}

// 无参数、单返回值
func (t Test) method3() (i int) {
	return
}

// 多参数、多返回值
func (t Test) method4(x, y int) (z int, err error) {
	return
}

// 无参数、无返回值
func (t *Test) method5() {

}

// 单参数、无返回值
func (t *Test) method6(i int) {

}

// 多参数、无返回值
func (t *Test) method7(x, y int) {

}

// 无参数、单返回值
func (t *Test) method8() (i int) {
	return
}

// 多参数、多返回值
func (t *Test) method9(x, y int) (z int, err error) {
	return
}

// 下面定义一个结构体类型和该类型的一个方法：
// 结构体
type User struct {
	Name  string
	Email string
}

// 方法
// 我们修改 Notify 方法，让它的接受者使用指针类型：
func (u *User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}
