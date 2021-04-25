go-example ![alt tag](https://api.travis-ci.org/phishman3579/java-algorithms-implementation.svg?branch=master)
==============================

Go Example project for Golang. 个人学习Go语言案例项目。



## Table of Contents
+ [Go简介](https://github.com/loveisontheway/go-example#Go简介)
+ [Go语言的主要特征](https://github.com/loveisontheway/go-example#Go语言的主要特征)
+ [Go适合做什么](https://github.com/loveisontheway/go-example#Go适合做什么)
+ [Go语言命名](https://github.com/loveisontheway/go-example#Go语言命名)
+ [Go语言声明](https://github.com/loveisontheway/go-example#Go语言声明)
+ [项目包说明](https://github.com/loveisontheway/go-example#项目包说明)
+ [main函数](https://github.com/loveisontheway/go-example#main函数)

## Go简介
    go语言（或 Golang）是Google开发的开源编程语言，诞生于2006年1月2日下午15点4分5秒，
    于2009年11月开源，2012年发布go稳定版。Go语言在多核并发上拥有原生的设计优势，
    Go语言从底层原生支持并发，无须第三方库、开发者的编程技巧和开发经验。
    go是非常年轻的一门语言，它的主要目标是“兼具Python 等动态语言的开发速度和C/C++等编译型语言的性能与安全性”
    很多公司，特别是中国的互联网公司，即将或者已经完成了使用 Go 语言改造旧系统的过程。
    经过 Go 语言重构的系统能使用更少的硬件资源获得更高的并发和I/O吞吐表现。充分挖掘硬件设备的潜力也满足当前精细化运营的市场大环境。
    Go语言的并发是基于 goroutine 的，goroutine 类似于线程，但并非线程。可以将 goroutine 理解为一种虚拟线程。
    Go 语言运行时会参与调度 goroutine，并将 goroutine 合理地分配到每个 CPU 中，最大限度地使用CPU性能。
    开启一个goroutine的消耗非常小（大约2KB的内存），你可以轻松创建数百万个goroutine。

## Go语言的主要特征
    1.自动立即回收。
    2.更丰富的内置类型。
    3.函数多返回值。
    4.错误处理。
    5.匿名函数和闭包。
    6.类型和接口。
    7.并发编程。
    8.反射。
    9.语言交互性。

## Go适合做什么
    服务端开发
    分布式系统，微服务
    网络编程
    区块链开发
    内存KV数据库，例如boltDB、levelDB
    云平台

## Go语言命名
    Go的函数、变量、常量、自定义类型、包(package)的命名方式遵循以下规则：
    1）首字符可以是任意的Unicode字符或者下划线
    2）剩余字符可以是Unicode字符、下划线、数字
    3）字符长度不限

## Go语言声明
    var（声明变量）
    const（声明常量）
    type（声明类型）
    func（声明函数）

## 项目包说明
| `package name` | `description` |
| :------ | :------ |
| basic | Go基础（变量、常量、数组、切片、指针、Map、结构体） |
| concurrency | Go并发编程（goroutine、runtime、channel、定时器、select、Sync等） |
| db | 数据操作（mysql、redis） |
| flow | 流程控制（if、switch、select、for、range） |
| function | 函数（参数、返回值、匿名函数、闭包、递归、defer、异常、测试） |
| gin | gin框架（gin路由、API\URL\表单参数、上传文件、JSON数据解析） |
| method | 方法（方法集、表达式、error） |
| network | 网络编程（TCP、UDP、HTTP） |
| object | 面向对象（匿名字段、接口） |
| test | 单元测试 |

## 基本类型介绍
| `类型` | `长度(字节)` | `默认值` | `说明` |
| :------ | :------ | :------ | :------ |
| bool | 1 | false |  |
| byte | 1 | 0 | uint8 |
| rune | 4 | 0 | Unicode Code Point, int32 |
| int, uint | 4或8 | 0 | 32 或 64 位 |
| int8, uint8 | 1 | 0 | -128 ~ 127, 0 ~ 255，byte是uint8 的别名 |
| int16, uint16 | 2 | 0 | -32768 ~ 32767, 0 ~ 65535 |
| int32, uint32 | 4 | 0 | -21亿~ 21亿, 0 ~ 42亿，rune是int32 的别名 |
| int64, uint64 | 8 | 0 | uint8 |
| float32 | 4 | 0 | 0.0 |
| float64 | 8 | 0 | 0.0 |
| complex64 | 8 |  |  |
| complex128 | 16 |  |  |
| uintptr | 4或8 |  | 以存储指针的 uint32 或 uint64 整数 |
| array |  |  | 值类型 |
| struct |  |  | 值类型 |
| string |  | "" | UTF-8 字符串 |
| slice |  | nil | 引用类型 |
| map |  | nil | 引用类型 |
| channel |  | nil | 引用类型 |
| interface |  | nil | 接口 |
| function |  | nil | 函数 |

## main函数
- Go语言程序的默认入口函数(主函数)：func main()
- 函数体用｛｝一对括号包裹。
```go
package main

func main() {
    // 函数体 
}
```