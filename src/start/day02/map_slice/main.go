package main

import "fmt"

// map和slice的结合
func main() {
	// 元素是map的切片 start

	//定义一个map类型的切片  需要对内部的map进行初始化   也就是这边需要两个make
	s1 := make([]map[int]string, 10)

	//之后 对内部的map进行初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "A"
	fmt.Println(s1)
	fmt.Println("-----------------------------------------------------------------")
	// 元素是map的切片 end

	//值为切片的map  首先要用make去初始化 申请内存空间 之后 map的key是string value是切片[]int
	m1 := make(map[string][]int, 10)
	m1["beijing"] = []int{10, 20, 30}
	fmt.Println(m1)
}
