package main

import "fmt"

//切片的复习
func main() {

	var s1 []int
	fmt.Println(s1)        // 因为没有分配内存  所以结果是[]等于空 nil  此时此刻没有分配内存也就是要进行初始化操作
	fmt.Println(s1 == nil) //true

	// 初始化方式1
	s1 = []int{1, 2, 3, 4}
	fmt.Println(s1) // [1 2 3 4]

	// 初始化方式2
	s2 := make([]bool, 10, 10)
	fmt.Println(s2) //[false false false false false false false false false false]

	// ---------------------------------------------------

	s3 := make([]int, 0, 4)
	fmt.Println(s3 == nil) //此时s3有分配内存了 所以是 false  也就是不为nil

	s1 = s1[0:2]    //也就是下标0 到下标1
	fmt.Println(s1) //[1 2]

	// ---------------------------------------------------
	// 切片的赋值拷贝
	sc1 := []int{1, 2, 3}
	sc2 := sc1
	sc2[1] = 100
	fmt.Println(sc1) //[1 100 3]  此时 是因为 切片不存值 而且 修改的值 对应的 是切片的底层数组的值 所以 两个结果都相同
	fmt.Println(sc2) //[1 100 3]

}
