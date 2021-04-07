package main

import "fmt"

// *** 方法集 ***
// Golang方法集 ：每个类型都有与之关联的方法集，这会影响到接口实现规则。
/*
   • 类型 T 方法集包含全部 receiver T 方法。
   • 类型 *T 方法集包含全部 receiver T + *T 方法。
   • 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
   • 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。
   • 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。
 */
// 用实例 value 和 pointer 调用方法 (含匿名字段) 不受方法集约束，
// 编译器总是查找全部方法，并自动转换 receiver 实参。
// Go 语言中内部类型方法集提升的规则：
// 类型 T 方法集包含全部 receiver T 方法。
func main()  {
	t1 := T{1}
	fmt.Printf("t1 is : %v\n", t1)
	t1.test()
}

type T struct {
	int
}

func (t T) test() {
	fmt.Println("类型 T 方法集包含全部 receiver T 方法。")
}
