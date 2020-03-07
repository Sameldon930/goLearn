package main

import "fmt"

// append添加元素
func main() {

	// 声明一个切片
	var s1 []int
	//往s1增加元素  此时的append 实现了 初始化切片 并给切片增加元素
	s1 = append(s1, 1)
	fmt.Println(s1)
}
