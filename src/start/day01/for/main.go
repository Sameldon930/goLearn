package main

import "fmt"

func main() {

	// 基本格式 0~9
	for i := 0; i < 10; i++ {
		// fmt.Println(i)
	}

	//写法2  5~9
	i := 5
	for ; i < 10; i++ {
		// fmt.Println(i)
	}

	// 写法3
	b := 6
	for b < 10 {
		// fmt.Println(b)
		b++
	}
	// for range 循环
	str := "hello张泽山"
	// i是key v是value str是要循环的数据  %c是输出字符  一个中文 3个key
	for i, v := range str {
		// fmt.Printf("%d  %c\n", i, v)
	}

	//九九乘法表

	//定义行的初始值
	for n := 1; n <= 9; n++ {
		//行有几列 第一行有一列 第二行有二列 以此类推
		for N := 1; N <= n; N++ {
			fmt.Printf("%d*%d=%d ", n, N, n*N)
		}
		//手动生成回车
		fmt.Println()

	}

}
