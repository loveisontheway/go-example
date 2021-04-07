package main

import "fmt"

// *** 条件语句switch ***
// switch 语句用于基于不同条件执行不同动作，
// 每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。
// Golang switch 分支表达式可以是任意类型，不限于常量。
// 可省略 break，默认自动终止。
/*
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
 */
func main()  {

	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70 : grade = "C"
		default: grade = "D"
	}

	switch {
		case grade == "A" :
			fmt.Printf("优秀!\n" )
		case grade == "B", grade == "C" :
			fmt.Printf("良好\n" )
		case grade == "D" :
			fmt.Printf("及格\n" )
		case grade == "F":
			fmt.Printf("不及格\n" )
		default:
			fmt.Printf("差\n" )
	}
	fmt.Printf("你的等级是 %s\n", grade )

	// Type Switch
	// switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	/*
	switch x.(type){
		case type:
		   statement(s)
		case type:
		   statement(s)
		// 你可以定义任意个数的case
		default: // 可选
			statement(s)
	}
	 */
	var x interface{}
	switch i := x.(type) { // 带初始化语句
	case nil:
		fmt.Printf(" x 的类型 :%T\r\n", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
