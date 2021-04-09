package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// *** Redis ***
// Redis 是完全开源免费的，遵守BSD协议，是一个高性能的key-value数据库。
//
// Redis 与其他 key - value 缓存产品有以下三个特点：
/*
   1. Redis支持数据的持久化，可以将内存中的数据保存在磁盘中，
      重启的时候可以再次加载进行使用。
   2. Redis不仅仅支持简单的key-value类型的数据，
      同时还提供string、list（链表）、set（集合）、hash表等数据结构的存储。
   3. Redis支持数据的备份，即master-slave模式的数据备份。
 */
// 性能极高 – Redis能读的速度是110000次/s,写的速度是81000次/s，单机能够达到15w qps，通常适合做缓存。
// 丰富的数据类型 – Redis支持二进制案例的 Strings, Lists, Hashes, Sets 及 Ordered Sets 数据类型操作。

// 使用第三方开源的redis库: github.com/garyburd/redigo/redis
// 命令行输入 ：
// 		go get github.com/garyburd/redigo/redis
func main()  {
	// 连接redis
	c, err := redis.Dial("tcp", "192.168.0.202:6379")
	// 连接成功，再进行密码认证
	err = c.Send("auth","tyb123")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	fmt.Println("redis conn success")
	defer c.Close()

	// String类型Set、Get操作
	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)
	fmt.Println("--------------------")

	// String批量操作
	_, err = c.Do("MSet", "abc", 100, "efg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}
	r2, err := redis.Ints(c.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	for _, v := range r2 {
		fmt.Println(v)
	}
	fmt.Println("--------------------")

	// 设置过期时间
	_, err = c.Do("expire", "hjk", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("--------------------")

	// List队列操作
	_, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}
	r3, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r3)
	fmt.Println("--------------------")

	// Hash表
	_, err = c.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r4, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r4)
	fmt.Println("--------------------")

}
