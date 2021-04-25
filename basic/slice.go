package main

import (
	"fmt"
)

// 全局
var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素

// *** 切片(动态数组) ***
// slice 并不是数组或数组指针
// 切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
// 切片的长度可以改变，因此，切片是一个可变的数组
// 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制
// cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
// 切片的定义：var 变量名 []类型，
// 比如 var str []string  var arr []int。
// 如果 slice == nil，那么 len、cap 结果都等于 0。
func main() {
	// 创建切片的各种方式
	// 1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("为空")
	} else {
		fmt.Println("不为空")
	}

	// 2. :=
	s2 := []int{}
	// 3. make
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	// 4. 初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	// 5. 从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)

	fmt.Printf("全局变量：arr %v\n", arr)
	fmt.Printf("全局变量：slice0 %v\n", slice0)
	fmt.Printf("全局变量：slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)
	fmt.Printf("-----------------------------------\n")

	// 局部
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr2[2:8]
	slice6 := arr2[0:6]         //可以简写为 slice := arr[:end]
	slice7 := arr2[5:10]        //可以简写为 slice := arr[start:]
	slice8 := arr2[0:len(arr)]  //slice := arr[:]
	slice9 := arr2[:len(arr)-1] //去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)

	fmt.Println("--------------------------------------")

	// 通过make创建切片
	slice10 := make([]int, 10)
	slice11 := make([]int, 10)
	slice12 := make([]int, 10, 10)
	fmt.Printf("make局部slice3 ：%v\n", slice10)
	fmt.Printf("make局部slice4 ：%v\n", slice11)
	fmt.Printf("make局部slice5 ：%v\n", slice12)
	fmt.Println("--------------------------------------")
	// 读写操作实际目标是底层数组，只需注意索引号的差别。
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 当slice中有两个冒号时，即slice[start:index:max],
	// 它的容量就是(max - start),
	// 实际引用的数组时从数组start索引开始到max索引为止，但不包括max索引处的元素。
	// 如果是一个冒号，则max默认为 len(data)-start
	// 长度: len()表示可见元素有几个（也即直接打印元素看到的元素个数），
	// 容量: cap()表示所有元素有几个
	// -- 容量的用处在哪？在与当你用 append 扩展长度时，
	// 如果新的长度小于容量，不会更换底层数组，否则，go 会新申请一个底层数组，
	// 拷贝这边的值过去，把原来的数组丢掉。也就是说，容量的用途是：
	// 在数据拷贝和内存申请的消耗与内存占用之间提供一个权衡。
	// -- 长度，则是为了帮助你限制切片可用成员的数量，提供边界查询的。
	// 所以用 make 申请好空间后，需要注意不要越界
	s := data[2:5:6]
	//s[0] += 100
	//s[1] += 200
	fmt.Println(len(data))
	fmt.Printf("s长度: %d，容量为: %d\n", s, cap(s))
	fmt.Println(s)
	fmt.Println(data)
	fmt.Println("--------------------------------------")

	// 可直接创建 slice 对象，自动分配底层数组。
	s10 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	fmt.Println(s10, len(s10), cap(s10))

	s11 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
	fmt.Println(s11, len(s11), cap(s11))

	s12 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s12, len(s12), cap(s12))
	fmt.Println("--------------------------------------")

	// 使用 make 动态创建slice，避免了数组必须用常量做长度的麻烦。
	// 还可用指针直接访问底层数组，退化成普通数组操作。
	s13 := []int{0, 1, 2, 3}
	p := &s13[2] // *int, 获取底层数组元素指针。
	*p += 100
	fmt.Println(s13)
	fmt.Println("--------------------------------------")

	// 至于 [][]T，是指元素类型为 []T 。
	data2 := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data2)
	fmt.Println("--------------------------------------")

	// 可直接修改 struct array/slice 成员。
	d := [5]struct {
		o int
	}{}
	s14 := d[:]
	d[1].o = 10
	s14[2].o = 20
	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0])
	fmt.Println("--------------------------------------")

	// 用append内置函数操作切片（切片追加）
	var a1 = []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a1)
	var b1 = []int{4, 5, 6}
	fmt.Printf("slice b : %v\n", b1)
	c1 := append(a1, b1...)
	fmt.Printf("slice c : %v\n", c1)
	d1 := append(c1, 7)
	fmt.Printf("slice d : %v\n", d1)
	e1 := append(d1, 8, 9, 10)
	fmt.Printf("slice e : %v\n", e1)
	fmt.Println("--------------------------------------")

	// append ：向 slice 尾部添加数据，返回新的 slice 对象。
	s15 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s15)
	s16 := append(s15, 1)
	fmt.Printf("%p\n", &s16)
	fmt.Println(s15, s16)
	fmt.Println("--------------------------------------")

	// 超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
	data3 := [...]int{0, 1, 2, 3, 4, 10: 0}
	s17 := data3[:2:3]
	fmt.Println(s17)
	s17 = append(s17, 100, 200)	// 一次 append 两个值，超出 s17.cap 限制。
	fmt.Println(s17, data3)   	// 重新分配底层数组，与原数组无关。
	fmt.Println(&s17[0], &data3[0]) 	// 比对底层数组起始指针。
	// 从输出结果可以看出，append 后的 s17 重新分配了底层数组，
	// 并复制数据。如果只追加一个值，则不会超过 s.cap 限制，也就不会重新分配。
	// 通常以 2 倍容量重新分配底层数组。在大批量添加数据时，
	// 建议一次性分配足够大的空间，以减少内存分配和数据复制开销。
	// 或初始化足够长的 len 属性，改用索引号进行操作。
	// 及时释放不再使用的 slice 对象，避免持有过期数组，造成 GC 无法回收。
	fmt.Println("--------------------------------------")


	// slice中cap重新分配规律
	s21 := make([]int, 0, 1)
	c21 := cap(s21)
	for i := 0; i < 50; i++ {
		s21 = append(s21, i)
		if n := cap(s21); n > c21 {
			fmt.Printf("cap: %d -> %d\n", c21, n)
			c21 = n
		}
	}
	fmt.Println("--------------------------------------")

	// 切片拷贝
	s31 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s31)
	s32 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s32)
	copy(s32, s31)
	fmt.Printf("copied slice s1 : %v\n", s31)
	fmt.Printf("copied slice s2 : %v\n", s32)
	s33 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s33)
	s33 = append(s33, s32...)
	fmt.Printf("appended slice s3 : %v\n", s33)
	s33 = append(s33, 4, 5, 6)
	fmt.Printf("last slice s3 : %v\n", s33)
	fmt.Println("--------------------------------------")

	// copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。
	// 两个 slice 可指向同一底层数组，允许元素区间重叠。
	data4 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array data : ", data4)
	s41 := data[8:]
	s42 := data[:5]
	fmt.Printf("slice s1 : %v\n", s41)
	fmt.Printf("slice s2 : %v\n", s42)
	copy(s42, s41)
	fmt.Printf("copied slice s1 : %v\n", s41)
	fmt.Printf("copied slice s2 : %v\n", s42)
	fmt.Println("last array data : ", data4)
	// 应及时将所需数据 copy 到较小的 slice，以便释放超大号底层数组内存。
	fmt.Println("--------------------------------------")

	// slice遍历
	data5 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data5[:]
	for index, value := range slice {
		fmt.Printf("inde : %v , value : %v\n", index, value)
	}
	fmt.Println("--------------------------------------")


	// 切片resize（调整大小）
	var a3 = []int{1, 3, 4, 5}
	fmt.Printf("slice a3 : %v , len(a3) : %v\n", a3, len(a3))
	b3 := a3[1:2]
	fmt.Printf("slice b3 : %v , len(b3) : %v\n", b3, len(b3))
	c3 := b3[0:3]
	fmt.Printf("slice c3 : %v , len(c3) : %v\n", c3, len(c3))
	fmt.Println("--------------------------------------")

	// 字符串和切片（string and slice）
	str := "hello world"
	s51 := str[0:5]
	fmt.Println(s51)
	s52 := str[6:]
	fmt.Println(s52)
	fmt.Println("--------------------------------------")

	// string本身是不可变的，因此要改变string中字符。需要如下操作： 英文字符串：
	str2 := "Hello world"
	s60 := []byte(str2) // 中文字符需要用[]rune(str)
	s60[6] = 'G'
	s60 = s60[:8]
	s60 = append(s60, '!')
	str2 = string(s60)
	fmt.Println(str2)
	fmt.Println("--------------------------------------")

	// 含有中文字符串
	str3 := "你好，世界！hello world！"
	s70 := []rune(str3)
	s70[3] = '够'
	s70[4] = '浪'
	s70[12] = 'g'
	s70 = s70[:14]
	str3 = string(s70)
	fmt.Println(str3)
	// golang slice data[:6:8] 两个冒号的理解
	// 常规slice , data[6:8]，从第6位到第8位（返回6， 7），长度len为2，
	// 最大可扩充长度cap为4（6-9）
	// 另一种写法： data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，
	// 长度len为6，最大扩充项cap设置为8
	// a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x
	slice2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d11 := slice2[6:8]
	fmt.Println(d11, len(d11), cap(d11))
	d12 := slice2[:6:8]
	fmt.Println(d12, len(d12), cap(d12))
	fmt.Println("--------------------------------------")
	// 数组or切片转字符串
	// strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)
}
