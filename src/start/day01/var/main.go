package main

//导入包
import "fmt"

// 变量
// 小驼峰命名
// 	可以声明在函数外部和内部
// 	同一个作用域中 不能重复声明同名字变量
// 非全局变量必须使用输出 否则无法编译（全局变量就不会）

// 简短变量声明(只能在函数中使用)
// 		s3 := "zzz"

//普通声明
// var name string
// var age int

// 多重赋值
// 		x,y := "a","c"
// 	匿名变量(使用多重赋值的时候 如果想要忽略某个值 相当于扔掉)
// 		不占用命名空间 不会分配内存 不存在重复声明
// 		x , _ := "zzzz"

//批量声明  一般用来声明全局
var (
	name string // 默认为 ""
	age  int    // 默认为 0
	isOk bool   //默认为false
)

func main() {

	//变量需先声明 在使用 然后才可以变量赋值
	//如果在函数内声明了 就要使用
	//在外部就可以不使用
	name = "zzs"
	age = 18
	isOk = false
	fmt.Println(age)            //输出并换行
	fmt.Print(isOk)             //直接输出
	fmt.Printf("name:%s", name) //格式化输出

	var s1 string = "zzzz"

	fmt.Println(s1)

	var s2 = "123"
	fmt.Println(s2)

	//短变量声明   只能在函数中使用
	s3 := 333
	fmt.Println(s3)
	x, y := "a", "c"
	fmt.Println(x)
	fmt.Println(y)

}
