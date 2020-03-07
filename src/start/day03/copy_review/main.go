package main

import "fmt"

// copy的复习
func main() {
	s1 := []int{1, 2, 3}
	s3 := make([]int, 4, 5)
	// copy给新的变量  长度 必须要大于等于原来的长度 否则无法实现copy
	copy(s3, s1)
	fmt.Println(s3)

}
