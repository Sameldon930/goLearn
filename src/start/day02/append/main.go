package main

import "fmt"

// append  切片添加元素

func main() {
	s1 := []string{"北京", "上海", "广州"}
	// s1[3] = "深圳" //这个写法会导致报错 因为超出了索引的范围
	fmt.Printf("s1=%v\n len(s1)=%d\n cap(s1)=%d\n", s1, len(s1), cap(s1)) //3 3
	fmt.Println("------------------------------------------------------------")

	//append往切片增加元素  处理之后 必须用之前的变量去接收返回值
	s1 = append(s1, "深圳")
	fmt.Printf("s1=%v\n len(s1)=%d\n cap(s1)=%d\n", s1, len(s1), cap(s1)) // 4  6
	fmt.Println("------------------------------------------------------------")

	s1 = append(s1, "杭州")
	fmt.Printf("s1=%v\n len(s1)=%d\n cap(s1)=%d\n", s1, len(s1), cap(s1)) //5 6
	fmt.Println("------------------------------------------------------------")

	//增加切片到切片中  ...表示把切片拆开
	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...)                                                // ... 表示ss拆开 单独拿出来 放到切片中
	fmt.Printf("s1=%v\n len(s1)=%d\n cap(s1)=%d\n", s1, len(s1), cap(s1)) //5 6
	fmt.Println("------------------------------------------------------------")

	// 去掉切片中的元素  去掉西安  不影响容量  这里的操作 类似于 将要去掉的元素的左边 拼接 右边
	ss = append(ss[0:1], ss[2:]...)
	fmt.Println(ss)               //[武汉 苏州]
	fmt.Println(cap(ss), len(ss)) //3  2
	fmt.Println("------------------------------------------------------------")

	//创建数组
	arr1 := [...]int{1, 3, 5, 7, 9}
	//转成切片
	sss1 := arr1[:]
	//删除切片中的元素
	sss1 = append(sss1[:2], sss1[3:]...)
	fmt.Println(sss1) //[1 3 7 9]
	fmt.Println(arr1) //[1 3 7 9 9]
	fmt.Printf("%p\n", sss1)
	fmt.Println("------------------------------------------------------------")

}
