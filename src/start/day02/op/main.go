package main

import "fmt"

// 运算符
// 算数运算符  + - * / %
// 关系运算符  == != > >= < <=
//逻辑运算符   && || ！
//位运算符   针对二进制     & | ^ << >>
//赋值运算符 = += -= *=  /= %= <<= >>= &= != ^=
func main() {
	var (
		a = 6
		b = 2
	)
	//算数运算符
	fmt.Println(a + b) //8
	fmt.Println(a - b) //4
	fmt.Println(a * b) //12
	fmt.Println(a / b) //3
	fmt.Println(a % b) //0

	//关系运算符 同种类型才能进行比较
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a < b)
	fmt.Println(a > b)

	//逻辑运算符
	if age := 12; age > 12 && age < 15 {
		fmt.Println("A")
	} else if age > 10 || age < 2 {
		fmt.Println("B")
	} else {
		fmt.Println("c")
	}

	//位运算符
	one := 5               //101
	two := 2               //010
	fmt.Println(one & two) //0  有0则0
	fmt.Println(one | two) //7  有1则1
	fmt.Println(^one)      //-6   相同为0 不同为1
	fmt.Println(one << 2)  //20   左移2位  101=>10100  2的四次方加上2的2次方
	fmt.Println(one >> 1)  //2    右移1位  101=>010

	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d * %d = %d  ", x, y, x*y)
		}
		fmt.Println()
	}

	one++
	fmt.Println(one)
}
