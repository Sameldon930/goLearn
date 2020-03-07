package main

import "fmt"

// if条件判断
func main() {

	// 单个条件
	age := 18
	if age > 18 {
		fmt.Println("大于18")
	} else {
		fmt.Println("小于18")
	}

	sex := 1
	// 多个条件
	if sex == 0 {
		fmt.Println("等于0")
	} else if sex == 1 {
		fmt.Println("等于1")
	} else {
		fmt.Println("等于2")
	}

	//临时变量运用到if中 在其他作用域 该变量就无法使用 如下
	if str := 123; str > 100 {
		fmt.Println("大于100")
	} else {
		fmt.Println("小于或等于100")
	}

	//此时打印上面的str变量 会出现未定义的报错
	// fmt.Println(str)  //undefined: str

}
