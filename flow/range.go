package main

import "fmt"

// *** 循环语句range ***
// Golang range类似迭代器操作，返回 (索引, 值) 或 (键, 值)。
// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
/*
for key, value := range oldMap {
    newMap[key] = value
}
 */
// 可忽略不想要的返回值，或 "_" 这个特殊变量。
func main()  {
	s := "abc"
	// 忽略 2nd value，支持 string/array/slice/map。
	for i := range s {
		println(s[i])
	}
	// 忽略 index。
	for _, c := range s {
		println(c)
	}
	// 忽略全部返回值，仅迭代。
	for range s {

	}

	m := map[string]int{"a": 1, "b": 2}
	// 返回 (key, value)。
	for k, v := range m {
		println(k, v)
	}
	fmt.Println("--------------------")

	// *注意，range 会复制对象。
	a := [3]int{0, 1, 2}

	for i, v := range a { // index、value 都是从复制品中取出。

		if i == 0 { // 在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
		}

		a[i] = v + 100 // 使用复制品中取出的 value 修改原数组。

	}
	fmt.Println(a) // 输出 [100, 101, 102]。
	fmt.Println("--------------------")

	// 建议改用引用类型，其底层数据不会被复制。
	str := []int{1, 2, 3, 4, 5}
	for i, v := range str { // 复制 struct slice { pointer, len, cap }。
		if i == 0 {
			str = str[:3]  // 对 slice 的修改，不会影响 range。
			str[2] = 100 // 对底层数据的修改。
		}
		println(i, v)
	}
	/*
	另外两种引用类型 map、channel 是指针包装，而不像 slice 是 struct。
	for 和 for range有什么区别?
		主要是使用场景不同
		for可以
		遍历array和slice
		遍历key为整型递增的map
		遍历string
	for range可以完成所有for可以做的事情，却能做到for不能做的，包括
		遍历key为string类型的map并同时获取key和value
		遍历channel
	 */

	// 循环控制Goto、Break、Continue
	// 循环控制语句可以控制循环体内语句的执行过程。
	// GO 语言支持以下几种循环控制语句：
	// 1.三个语句都可以配合标签(label)使用
	// 2.标签名区分大小写，定以后若不使用会造成编译错误
	// 3.continue、break配合标签(label)可用于多层循环跳出
	// 4.goto是调整执行位置，与continue、break配合标签(label)的结果并不相同

}
