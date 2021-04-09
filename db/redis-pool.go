package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// *** Redis连接池 ***
func main()  {
	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池
	c.Send("auth","tyb123")

	_,err := c.Do("Set","abc",200)
	if err != nil {
		fmt.Println(err)
		return
	}

	r,err := redis.Int(c.Do("Get","abc"))
	if err != nil {
		fmt.Println("get abc faild :",err)
		return
	}
	fmt.Println(r)
	pool.Close() //关闭连接池
}

var pool *redis.Pool  //创建redis连接池

func init(){
	pool = &redis.Pool{     //实例化一个连接池
		MaxIdle:16,    //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:0,    //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout:300,    //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn ,error){     //要连接的redis数据库
			return redis.Dial("tcp","192.168.0.202:6379")
		},
	}
}
