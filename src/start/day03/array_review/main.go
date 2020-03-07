package main

import "fmt"

// 数组的复习
func main() {

	var a3 [10]int                 //声明
	a3 = [10]int{1, 2, 3, 4, 5, 6} //赋值
	a4 := [...]int{1, 2, 3, 4}     //声明并初始化
	a5 := [...]int{1: 100, 2: 200}

	fmt.Println(a3) //[1 2 3 4 5 6 0 0 0 0]
	fmt.Println(a4) //[1 2 3 4]
	fmt.Println(a5) //[0 100 200]

	// 多维数组的声明
	var aa [3][2]int //声明  有三个元素 每个元素内有两个元素
	fmt.Println(aa)

	//多维数组声明并初始化  只有最外层可以 ...
	aa2 := [...][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(aa2)

	a6 := [3]int{1, 2, 3}
	f1(a6)
	fmt.Println(a6) //[1 2 3]

	//数组的赋值特性
	a7 := [3]int{1, 2, 3}
	a8 := a7
	a8[1] = 11
	fmt.Println(a7) //[1 2 3]
	fmt.Println(a8) //[1 11 3]

}
func f1(x [3]int) {
	x[1] = 100 //此处修改的是副本的元素 而不是原来a6的值
}
