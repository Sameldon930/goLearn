package main

import "fmt"

// copy的用法 复制切片

func main() {
	a1 := []int{1, 2, 3}
	a2 := a1
	//定义新的一个切片 有长度的切片
	a3 := make([]int, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3) //[1 2 3] [1 2 3] [1 2 3]
	a1[0] = 100
	fmt.Println(a1, a2, a3) //[100 2 3] [100 2 3] [1 2 3]   因为a3已经在之前复制了a1 所以不变

	//去掉2
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1, len(a1), cap(a1))
}
