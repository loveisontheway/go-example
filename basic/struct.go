package main

import "fmt"

// *** 结构体 ***
// Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
// Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
// 多用于定义程序运行期间不会改变的那些值
func main()  {

	// 类型定义和类型别名的区别
	// 类型别名与类型定义表面上看只有一个等号的差异，
	// 我们通过下面的这段代码来理解它们之间的区别。
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
	// 结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。
	// b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。
	fmt.Println("-----------------------")

	// 结构体的定义
	/*
	type 类型名 struct {
		字段名 字段类型
		字段名 字段类型
		…
	}
	*/
	// 其中：
	//    1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
	//    2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
	//    3.字段类型：表示结构体字段的具体类型。

	// 结构体实例化
	// 只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
	// 结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
	//   var 结构体实例 结构体类型
	var p1 person
	p1.name = "pprof.cn"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={pprof.cn 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"pprof.cn", city:"北京", age:18}
	fmt.Println("-----------------------")

	// 匿名结构体
	// 在定义一些临时数据结构等场景下还可以使用匿名结构体。
	var user struct{Name string; Age int}
	user.Name = "pprof.cn"
	user.Age = 18
	fmt.Printf("%#v\n", user)
	fmt.Println("-----------------------")

	// 结构体初始化
	var p4 person
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
	fmt.Println("-----------------------")

	// 使用键值对初始化
	// 使用键值对对结构体进行初始化时，
	// 键对应结构体的字段，值对应该字段的初始值
	p5 := person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"pprof.cn", city:"北京", age:18}
	fmt.Println("-----------------------")

	// 也可以对结构体指针进行键值对初始化，例如：
	p6 := &person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"pprof.cn", city:"北京", age:18}
	fmt.Println("-----------------------")

	// 当某些字段没有初始值的时候，该字段可以不写。
	// 此时，没有指定初始值的字段的值就是该字段类型的零值。
	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}
	fmt.Println("-----------------------")

	// 使用值的列表初始化
	// 初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
	p8 := &person{
		"pprof.cn",
		"北京",
		18,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"pprof.cn", city:"北京", age:18}
	// 使用这种格式初始化时，需要注意：
	// 1.必须初始化结构体的所有字段。
	// 2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 3.该方式不能和键值初始化方式混用。
	fmt.Println("-----------------------")

	// 结构体内存布局
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
	fmt.Println("-----------------------")

	// 调用构造函数
	p9 := newPerson("pprof.cn", "测试", 90)
	fmt.Printf("%#v\n", p9)
	fmt.Println("-----------------------")

	// 嵌套结构体
	// 一个结构体中可以嵌套包含另一个结构体或结构体指针。
	user1 := User{
		Name:   "pprof",
		Gender: "女",
		Address: Address{
			Province: "黑龙江",
			City:     "哈尔滨",
		},
	}
	fmt.Printf("user1=%#v\n", user1)//user1=main.User{Name:"pprof", Gender:"女", Address:main.Address{Province:"黑龙江", City:"哈尔滨"}}
	fmt.Println("-----------------------")

	// 结构体的“继承”
	// Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
	d1 := &Dog {
		Feet: 4,
		Animal: &Animal {	// 注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.move()
	d1.wang()
	fmt.Println("-----------------------")
}

// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

// Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

// 构造函数
// Go语言的结构体没有构造函数，我们可以自己实现。
// 例如，下方的代码就实现了一个person的构造函数。
// 因为struct是值类型，如果结构体比较复杂的话，
// 值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// 定义一个Person（人）结构体
type person struct {
	name string
	city string
	age  int8
	// 同样类型的字段也可以写在一行
	//name, city string
	//age        int8
}

//类型定义
type NewInt int

//类型别名
type MyInt = int
