package main

import "fmt"

// 函数 main函数没有返回值
// 函数的定义
// 参数没有默认值

// 封装函数  （参数）（返回值）
func sum(x int, y int) (ret int) {
	ret = x + y
	return
}

//没有返回值的
func f1() {
	fmt.Println(111)
}
func f11() {
	fmt.Println(1111)
}

//没有参数的 没有返回值 调用这种函数 就不能进行赋值了
func f2() {
	fmt.Println("f2222222")
}
func f22() {
	fmt.Println("f22333222")
}

//没有参数 有返回值的
func f33() int {
	return 3
}
func f3() int {
	return 3
}

// 返回值可以命名 也可以不命名

// 参数没有命名 也可以命名 下面的ret就相当于已经定义好了
func f4(x int, y int) (ret int) {
	ret = x + y
	return
}
func f44(x int, y int) (ret int) {
	ret = x + y
	return
}

//多个返回值
func f5() (int, string) {
	return 1, "zzs"
}
func f55() (int, string) {
	return 1, "zzs"

}

//参数类型相同的情况下简写 此时 x和y都是int  以此类推
func f6(x, y int, m, n string, i, j bool) {

}

// 可变长参数也就是可以传多个但是其实就是一个切片的元素  必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x) //zzs
	fmt.Println(y) //[1 2 3 4 5]  int类型的切片

}
func main() {
	// 调用函数
	r := sum(1, 2)
	println(r)

	//调用多个返回值的函数 进行赋值
	a, b := f5()
	println(a)
	println(b)
	fmt.Println("--------------------------")

	f2()
	f3 := f3()
	fmt.Println(f3)
	fmt.Println("--------------------------")

	f7("zzs", 1, 2, 3, 4, 5, 6)
	// 调用可变长的参数的函数
	f7("zzs", 1, 2, 3, 4, 5)

}
