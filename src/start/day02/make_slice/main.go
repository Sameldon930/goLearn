package main

import "fmt"

// make创建切片
// 切片的本质就是 数组底层的封装 属于引用类型 本身就是一个框 真正的数据都是保存在底层数组里

func main() {
	//参数1 元素类型  参数2 长度 参数3容量 如果没有填写第三个参数 则默认值等于第二个参数
	s1 := make([]int, 5, 10)                                            //如果长度为0  则输出没有值
	fmt.Printf("s1=%v\nlen(s1)=%d\ncap(s1)=%d\n", s1, len(s1), cap(s1)) //[0 0 0 0 0] 5 10

	// 切片之间不能进行比较  只能判断是否为空 并且必须要用 len(s1)== 0 来判断  不能使用 s1 == nil

	// 一个nil值的切片并没有底层数组
	// 一个nil值的切片的长度和容量都是0
	// 一个长度和容量都是0的不一定是nil

	// 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println(s3, s4) //[1 3 5] [1 3 5]
	s3[0] = 100
	fmt.Println(s3, s4) //[100 3 5] [100 3 5]

	// 切片的循环
	// 1.索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(i)
	}
	fmt.Println("----------------------")
	//for range
	for k, _ := range s3 {
		fmt.Println(k)
	}
	fmt.Println("----------------------")

}
