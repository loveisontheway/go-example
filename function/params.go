package main

import "fmt"

// *** 函数参数 ***
// 函数定义时指出，函数定义时有参数，该变量可称为函数的形参。
// 形参就像定义在函数体内的局部变量。
// 但当调用函数，传递过来的变量就是函数的实参，函数可以通过两种方式来传递参数：
// 值传递：指在调用函数时将实际参数复制一份传递到函数中，
// 这样在函数中如果对参数进行修改，将不会影响到实际参数。
/*
	func swap(x, y int) int {
		   ... ...
	}
 */
func main()  {
	var a, b int = 1, 2
	/*
	   调用 swap() 函数
	   &a 指向 a 指针，a 变量的地址
	   &b 指向 b 指针，b 变量的地址
	*/
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println("-------------------")

	// 使用 slice 对象做变参时，必须展开。（slice...）
	s := []int{1, 2, 3}
	res := testSlice("sum: %d", s...)    // slice... 展开slice
	println(res)

}

/* 定义相互交换值的函数 */
func swap(x, y *int) {
	var temp int
	temp = *x /* 保存 x 的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y*/
}

/*
func myfunc(args ...int) {    //0个或多个参数
}

func add(a int, args…int) int {    //1个或多个参数
}

func addMore(a int, b int, args…int) int {    //2个或多个参数
}
*/

func testSlice(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}