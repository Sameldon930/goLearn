package main

import "fmt"

// 切片  和数组的区别就是在于 长度是可变的  数组的长度是固定的 所以定义切片的时候 长度没有填写固定
// len() 求出长度
// cap()求出切片的容量
// 容量是指底层数组的容量 也就是 切片的第一个元素 在底层数组算起 直到最后一个 加起来的长度 就是容量
// 比如  数组 1 2 3 4 5  切片 [2 3 4] 那么 这个切片长度为3 容量为4（2在数组中开始算起到结束 一共四个 所以容量为4）

func main() {

	var s1 []int        //定义整型切片
	var s2 []string     //定义字符串切片
	fmt.Println(s1, s2) // [] []

	//判断是否为空 nil
	fmt.Println(s1 == nil) //true
	fmt.Println(s2 == nil) //true

	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"张", "泽", "山"}
	fmt.Println(s1) //[1 2 3]
	fmt.Println(s2) //[张 泽 山]

	fmt.Println(len(s1), cap(s1)) // 3 3
	fmt.Println(len(s2), cap(s2)) // 3 3

	//由数组得到切片 基于数组切割成数组  左闭右开  例如下面下标到4 但是切割的时候到4 是不包含下标为4的 以此类推
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4]
	fmt.Println(s3) //[1 3 5 7]
	s4 := a1[1:6]   //也就是 下标为1到下标为5  [3 5 7 9 11]
	s5 := a1[:4]    //下标从0到3[1 3 5 7]
	s6 := a1[3:]    //下标从3到最后一个[7 9 11 13]
	s7 := a1[:]     //从0到最后一个[1 3 5 7 9 11 13]
	fmt.Println(s4, s5, s6, s7)

	//s5的长度和容量
	fmt.Printf("len(s5):%d,cap(s5):%d \n", len(s5), cap(s5)) //4 7

	// 切片再切片
	s8 := s6[1:2]
	fmt.Println(s8)                                        //[9]
	fmt.Printf("len(s8):%d, cap(s8):%d", len(s8), cap(s8)) //1  3

}