package main

import "fmt"

// 数组练习题

func main() {

	//1. 求出数组 [1 3 5 7 8]所有元素的和
	ex1 := [5]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range ex1 {
		sum += v
	}
	fmt.Println(sum)
	fmt.Println("----------------------")

	//2.找出数组中下标的对应的值加起来等于8的下标  如 [1 2] [0 3]
	ex2 := [5]int{1, 3, 5, 7, 8}

	for i := 0; i < len(ex2); i++ {
		//第二循环的key比上一层加1  当i=1 j=2 以此类推
		for j := i + 1; j < len(ex2); j++ {
			//当 两层key的值等于8 那么就输出
			if ex2[i]+ex2[j] == 8 {
				fmt.Printf("[%d %d] ", i, j)
			}
		}
	}
}
