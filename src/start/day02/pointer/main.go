package main

import "fmt"

// 指针 记住两个符号   16进制
// 1、&取出地址
// 2、*根据地址取出地址指向的值

// make和new都是用来申请内存的
// new很少使用 一般都是用来给基本类型申请内存的  如 int string  返回的是 对应类型的指针
// make 用来给slice map channel 申请内存  返回的是对应的类型的本身

func main() {

	n := 18
	//取地址
	fmt.Println(&n) //0xc00006e068
	// 打印出类型
	fmt.Printf("%T\n", n) //int
	//根据地址 取值
	fmt.Println(*&n)

	b := &n        //n地址 赋值给 b
	fmt.Println(b) //0xc00006e068
	//打印出b的地址
	fmt.Println(&b) //0xc000098020

	// var a *int //此时初始化的int 没有值 也就是nil 所以 没有对应的地址的值
	// *a = 100   //100赋值给a的地址对应的值
	// fmt.Println(*a) //报错

	// 申请一个内存地址
	var a = new(int)
	fmt.Println(a)  //0xc0000120f8
	fmt.Println(*a) //地址对应的值  0
	*a = 100
	fmt.Println(*a) //赋值之后  100

}
