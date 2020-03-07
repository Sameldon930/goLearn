package main

import "fmt"

//切片的练习
func main() {

	a := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a) //[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]

	//练习2
	a1 := [...]int{1, 2312, 434, 545, 46, 575}
	s1 := a1[:]
	//去掉索引为1的数字
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1) //[1 434 545 46 575]
}
