package main

import "fmt"

// 整型
// 分为两大类
// 			1.无符号：uint8 uint16 uint32
// 			2.有符号：int8 int16 int32 int64
// 		  uint在32位操作系统上是uint32  63位上是uint64
// 			int在32位操作系统上是int32  63位上是int64
// 			uintptr 无符号整型 用于存放一个指针

// 		Go无法直接定义二进制数

func main() {
	//十进制
	i1 := 101
	fmt.Printf("%d\n", i1) //直接输出   101
	fmt.Printf("%b\n", i1) //十进制转二进制   1100101
	fmt.Printf("%o\n", i1) //十进制转八进制   145    ---权限
	fmt.Printf("%x\n", i1) //十进制转十六进制   65     ---内存
	fmt.Printf("%X\n", i1) //十进制转十六进制   65     ---内存

	//八进制
	i2 := 077
	fmt.Printf("%b\n", i2) //转换二进制     111111
	fmt.Printf("%d\n", i2) //转换十进制     63
	fmt.Printf("%x\n", i2) //转换十六进制   3f

	//十六进制
	i3 := 0x1234567
	fmt.Printf("%b\n", i3) //直接输出二进制    1001000110100010101100111
	fmt.Printf("%d\n", i3) //直接输出十进制    19088743
	fmt.Printf("%o\n", i3) //直接输出八进制	   110642547

	// 声明int8类型的整型
	i4 := int8(9)
	fmt.Printf("%d\n", i4) //十进制 9

	//输出变量类型
	fmt.Printf("%T\n", i4) //int8
}
