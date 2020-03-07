package main

import "fmt"

// 布尔型
// 默认为false
// 			禁止整型强制转换成布尔型
// 			无法参与数值运算 也无法与其他类型转换
func main() {
	b1 := false
	var b2 bool
	fmt.Printf("%v\n", b1) //打印值false
	fmt.Printf("%v\n", b2) //打印出来的是 布尔型的默认值  false

}
