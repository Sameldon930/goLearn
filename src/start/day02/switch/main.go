package main

import "fmt"

func main() {

	m := 5
	switch m {
	case 1:
		fmt.Println("a")
		break
	case 2:
		fmt.Println("b")
	case 3:
		fmt.Println("c")
	case 4:
		fmt.Println("c")
	default:
		fmt.Println("Z")
	}

	//临时变量的写法

	switch n := 5; n {
	case 1:
		fmt.Println("a")
		break
	case 2:
		fmt.Println("b")
	case 3:
		fmt.Println("c")
	case 4:
		fmt.Println("c")
	default:
		fmt.Println("Z")
	}

	//case 多个值的写法
	switch t := 5; t {
	case 1, 3, 5, 7, 9:
		fmt.Println("a")
		break
	case 2, 4, 6, 8:
		fmt.Println("b")
	default:
		fmt.Println("Z")
	}

	//判断写法
	p := 5
	switch {
	case p < 5:
		fmt.Println("a")
	case p > 25 && p < 35:
		fmt.Println("b")
	default:
		fmt.Println("Z")
	}

	//u=5的时候 跳过
	for u := 1; u < 10; u++ {
		if u == 5 {
			continue
		}
		fmt.Println(u)
	}
}
