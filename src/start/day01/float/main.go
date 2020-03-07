package main

import "fmt"

// 浮点数
// Go语言中默认的小数都是64位
func main() {
	//32位 浮点数 最大值
	// math.MaxFloat32
	//64位 浮点数 最大值
	// math.MaxFloat64

	f1 := 1.23456 //默认64位
	fmt.Printf("%f\n", f1)
	fmt.Printf("%T \n ", f1) //float64

	//强制转换成32位
	f2 := float32(1.23456)
	fmt.Printf("%f \n", f2)
	fmt.Printf("%T \n", f2)

}
