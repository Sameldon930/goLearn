package main

// 字符串
// 字符串是用双引号包裹的    "hello world"
// 			单引号包裹的是字符（字符就是一个字母 汉字 符号）  'H'  字符通常是int32  一个字符等于一个字节

// 以下符号需要转义
// 			\r 回车符 \n换行符
// 			\t 制表符 \' 单引号
// 			\" 双引号 \\斜杠

import (
	"fmt"
	"strings"
)

func main() {

	// D:\goLearn\src\start 输出这个地址  \需要转义
	str := "D:\\goLearn\\src\\start"
	fmt.Printf("%s \n", str)

	str2 := "i 'm ok"
	fmt.Println(str2)
	// 多行字符串
	str3 := `
		world
	`
	fmt.Printf("%s\n", str3)

	//字符串长度
	fmt.Println(len(str)) //20

	//字符串拼接
	name := "learn"
	world := "Go"
	fmt.Println(name + world)                 //方法1  learnGo
	fmt.Printf("%s%s\n", name, world)         //方法2  learnGo
	third := fmt.Sprintf("%s%s", name, world) //方法3
	fmt.Println(third)                        //learnGo
	// 方法2和方法3的区别在于  printf是直接输出  Sprintf是需要一个变量进行接收

	// 分割  需要import strings进来
	sp := strings.Split(str, "\\")
	fmt.Println(sp) //[D: goLearn src start]

	//是否包含 返回布尔型  需要import strings进来
	sc := strings.Contains(str, "s")
	fmt.Println(sc) //true

	// 前缀  需要import strings进来  返回布尔型
	fmt.Println(strings.HasPrefix(str, "D"))
	// 后缀  需要import strings进来  返回布尔型
	fmt.Println(strings.HasSuffix(str, "t"))

	newStr := "abcdef"
	//字符串出现的位置 从左到右0为初始  需要import strings进来
	fmt.Println(strings.Index(newStr, "c")) //2
	//字符串出现的位置  从右到左1位初始 需要import strings进来
	fmt.Println(strings.LastIndex(newStr, "d")) //3

	//join 用某个字符串连接
	fmt.Println(strings.Join(sp, "-"))

	//字符串的修改
	u1 := "black"
	u2 := []rune(u1)        //rune 将字符串u1转成切片 才能对字符进行修改
	u2[0] = 'a'             //修改第一个字符 所以这边是 单引号
	fmt.Println(string(u2)) //将切片转成字符串并输出  alack

	u3 := "红色"
	fmt.Println([]rune(u3)) //切片的元素值是 string

	cc1 := "红"
	cc2 := '红'
	fmt.Printf("%T\n", cc1) //string
	fmt.Printf("%T\n", cc2) //int32

	cc3 := "H"
	cc4 := 'H'
	fmt.Printf("%T\n", cc3) //string
	fmt.Printf("%T\n", cc4) //string

	//类型转换  只有整型和浮点型能转换 还有字符串和切片
	num := 10
	var f1 float64
	f1 = float64(num)
	fmt.Println(f1)        //10
	fmt.Printf("%T\n", f1) //float64
}
