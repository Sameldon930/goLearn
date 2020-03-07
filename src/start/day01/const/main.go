package main

import "fmt"

// 常量  const

// 常量计数器 iota
// 		    只能在常量的表达式中使用
// 			iota在const关键字出现 将被重置为0
// 			每新增行一行const的声明都会将计数加1
// 			能够简化定义 在定义枚举的时候有用

const pi = 123

const (
	statusOk   = 200
	statusFail = 404
)

//批量声明
const (
	a = 1
	b
	c
)

//iota 从01开始 后续自动加1
const (
	n1 = iota //0
	n2        //1
	n3        //2
	n4        //3
)

const (
	b1 = iota //0
	b2        //1
	_         //2 但是被弃用
	b3        //3
)
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //隔了一行变量声明 所以此时为2
	c4 = iota //同理为3
)

const (
	d1, d2 = iota + 1, iota + 2 //1 2
	d3, d4 = iota + 1, iota + 2 //2 3
)

func main() {
	fmt.Println(d1, d2, d3, d4)

}
