package main

import "fmt"

// *** 表达式 ***
// Golang 表达式 ：根据调用者不同，方法分为两种表现形式:
/*
	instance.method(args...) ---> <type>.func(instance, args...)
 */
// 前者称为 method value，后者 method expression。
//
// 两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，
// 而 method expression 则须显式传参。
func main()  {
	u := User2{1, "Tom"}
	u.Test()

	mValue := u.Test
	mValue() // 隐式传递 receiver

	mExpression := (*User2).Test
	mExpression(&u) // 显式传递 receiver
}

type User2 struct {
	id   int
	name string
}

func (self *User2) Test() {
	fmt.Printf("%p, %v\n", self, self)
}