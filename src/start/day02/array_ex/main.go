package main

import "fmt"

// 数组练习题

func main() {

	//1. 求出数组 [1 3 5 7 8]所有元素的和
	q1 := [...]int{1, 3, 5, 7, 8}
	//需要定义在外部 才能在循环体外部输出
	sum := 0
	for _, v := range q1 {
		sum += v
	}
	fmt.Println(sum)

	fmt.Println("----------------------")

	//2.找出数组中下标的对应的值加起来等于8的下标  如 [1 2] [0 3]
	q2 := [...]int{1, 3, 5, 7, 8}

	//第一列
	for i := 0; i < len(q2); i++ {
		//第二列
		for j := i + 1; j < len(q1); j++ {
			if q2[i]+q2[j] == 8 {
				fmt.Printf("[%d %d]", i, j)
			}
		}
	}

}
