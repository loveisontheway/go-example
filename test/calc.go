package main

import "testing"

// *** 单元测试 ***
// Go语言中的测试依赖go test命令。
// 编写测试代码和编写普通的Go代码过程是类似的，
// 并不需要学习新的语法、规则或工具。
/*
类型	    格式	                作用
测试函数	函数名前缀为Test	    测试程序的一些逻辑行为是否正确
基准函数	函数名前缀为Benchmark	测试函数的性能
示例函数	函数名前缀为Example	为文档提供示例文档
 */
// go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，
// 然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、
// 报告测试结果，最后清理测试中生成的临时文件。
// Golang单元测试对文件名和方法名，参数都有很严格的要求。
/*
   1、文件名必须以xx_test.go命名
   2、方法必须是Test[^a-z]开头
   3、方法参数必须 t *testing.T
   4、使用go test执行单元测试
 */
func main()  {
	
}

//  测试函数的格式
// 每个测试函数必须导入testing包，测试函数的基本格式（签名）如下：
func TestName(t *testing.T){
	// ...
}

// 测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头
func TestAdd(t *testing.T){ /*...*/ }
func TestSum(t *testing.T){ /*...*/ }
func TestLog(t *testing.T){ /*...*/ }