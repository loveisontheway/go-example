package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 全局: 一维数组
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

// 全局: 多维数组
var arr3 [5][3]int
var arr4 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

// *** 数组 ***
// 数组：是同一种数据类型的固定长度的序列。
// 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
// 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
// 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
func main() {

	// 局部: 一维数组
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度
	c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 18}, // 可省略元素类型
		{"user2", 20}, // 别忘了最后一行的逗号
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
	fmt.Println("------------------")
	// 局部: 多维数组
	m := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	n := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."
	fmt.Println(arr3, arr4)
	fmt.Println(m, n)
	fmt.Println("------------------")

	// 值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。
	num := [2]int{}
	fmt.Printf("num: %p\n", &num)
	test(num)
	fmt.Println(num)

	// 内置函数 len 和 cap 都返回数组长度 (元素数量)。
	println(len(num), cap(num))

	// 遍历多维数组
	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}

	fmt.Println("----------------")
	// 数组拷贝和传参
	var arr6 [5]int
	printArr(&arr6)
	fmt.Println(arr6)
	fmt.Println("----------------")
	arr7 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr7)
	fmt.Println(arr7)
	fmt.Println("----------------")

	// 求数组所有元素之和
	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	// rand.Seed(1)
	rand.Seed(time.Now().Unix())

	var z [10]int
	for i := 0; i < len(z); i++ {
		// 产生一个0到1000的随机数
		z[i] = rand.Intn(1000)
	}
	sum := sumArr(z)
	fmt.Println(z)
	fmt.Printf("sum=%d\n", sum)

	// 找出数组中和为给定值的两个元素的下标，
	// 例如数组[1,3,5,8,7]，找出两个元素之和等于8的下标分别是（0，4）和（1，2）
	arr8 := [5]int{1, 3, 5, 8, 7}
	fmt.Println("---------------")
	myArr(arr8, 8)
}

// 求元素和，按指定下标的值（求两值相加等于8），target = 8
func myArr(num [5]int, target int) {
	// 遍历数组
	for i := 0; i < len(num); i++ {
		other := target - num[i]
		// 再次遍历
		for j := i + 1; j < len(num); j++ {
			if num[j] == other {
				fmt.Printf("(%d,%d)%d,%d\n", i, j, num[i], num[j])
			}
		}
	}
}

// 求元素之和
func sumArr(num [10]int) int {
	var sum int = 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}
