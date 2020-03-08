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
		fmt.Printf("%d  %c\n", i, v)
	}

	//九九乘法表

	for n := 1; n <= 9; n++ {
		for m := 1; m <= n; m++ {
			fmt.Printf("%d*%d = %d  ", n, m, n*m)
		}
		fmt.Println()
	}

}
