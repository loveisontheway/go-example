package main

import "fmt"

// *** 接口 ***
// 接口（interface）定义了一个对象的行为规范，
// 只定义规范不实现，由具体的对象来实现规范的细节。，
// 在Go语言中接口（interface）是一种类型，一种抽象的类型。
// interface是一组method的集合，是duck-type programming的一种体现。
// 接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，
// 我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。
// 请牢记接口（interface）是一种类型。
// 每个接口由数个方法组成，接口的定义格式如下：
// 	  type 接口类型名 interface{
//        方法名1( 参数列表1 ) 返回值列表1
//        方法名2( 参数列表2 ) 返回值列表2
//        …
//    }
// 其中：
// 1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，
//   一般会在单词后面添加er，如有写操作的接口叫Writer，
//   有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
// 2.方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，
//   这个方法可以被接口所在的包（package）之外的代码访问。
// 3.参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。
func main()  {
	// 一个类型实现多个接口
	// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。
	// 例如，狗可以叫，也可以动。我们就分别定义Sayer接口和Mover接口，如下： Mover接口。
	var x Sayer
	var y Mover

	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()

	var b = car{brand: "保时捷"}
	y = b
	y.move()
	fmt.Println("----------------------")

	// 多个类型实现同一接口

}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

type car struct {
	brand string
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 当你看到这个接口类型的值时，你不知道它是什么，
// 唯一知道的就是可以通过它的Write方法来做一些事情。
type writer interface{
	Write([]byte) error
}

