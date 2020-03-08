package main

import "fmt"

//数组的定义 必须要声明他的长度（具体数字 或者 ...） 没有长度的 是切片
// var arr [长度]存放的数据类型

// 数组  var a [3]int  定义 名字为a 长度为3 存放的元素类型int的数组
// 设置数组的初始值[3]bool{true,true,true}
//长度或者类型不一致 无法进行比较
//长度是类型数组的一部分
func main() {
	var arr1 [3]bool
	var arr2 [4]bool
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)

	// 如果不进行初始化 那么所有的元素默认是都零 也就是false  整型和浮点型都是0 字符串为""
	fmt.Println(arr1, arr2) //[false false false] [false false false false]

	// 数组的初始化  用花括号指定所的规定的元素值
	// 数组的初始化 1 重新赋值(存放的类型和长度要一一致) 初始化 arr1
	arr1 = [3]bool{true, true, true}
	fmt.Println(arr1) //[true true true]

	//数组的初始化  2  根据初始值(存放的类型和长度要一一致) 数组长度自行生成
	arr2 = [...]bool{true, true, true, true}
	fmt.Println(arr2) //[true true true true]

	//数组的初始化3 定义key
	arr3 := [5]int{2: 1, 4: 2} //这样就表示 下标2的是1 下标4的是2  其他都是0
	fmt.Println(arr3)          //[0 0 1 0 2]
	fmt.Println("--------------------------")

	//数组的遍历
	cities := [...]string{"北京", "上海", "广州", "深圳"}

	//1.根据索引进行遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	}
	fmt.Println("--------------------------")
	//2 for range 遍历
	for i, v := range cities {
		fmt.Println(i, v)
	}

	//多维数组 如 [ [1 2] [3 4] [5 6]]
	aar := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(aar) //[ [1 2] [3 4] [5 6]]

	fmt.Println("--------------------------")

	// 多维数组的遍历
	for _, v1 := range aar {
		for _, v2 := range v1 {
			fmt.Println(v2) //1 2 3 4 5 6
		}
	}
	fmt.Println("--------------------------")

	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2) //b1 [1 2 3]   b2 [100 1 2]

}
