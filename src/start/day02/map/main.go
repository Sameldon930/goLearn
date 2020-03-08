package main

// map是一种无序的 key-value的数据结构
// 是引用类型
// 必须初始化才能使用
import (
	"fmt"
)

// 用法  var m1 map[key]value   如  var m1 map[string]int  表示这个map的key是string类型 value是int类型
func main() {

	//map初始化
	//通过make申请内存空间  申请的容量要考虑充分 避免在运行过程中再次扩容
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["zzs"] = 18
	m1["hh"] = 2
	fmt.Println(m1) //map[hh:2 zzs:18]
	fmt.Println("---------------------------------------------------")

	//判断map是否存在某个元素  ok 约定成俗为作为判断的key
	value, ok := m1["ooo"]
	if !ok {
		fmt.Println("无此元素")
	} else {
		fmt.Println(value)
	}
	// 如果直接输出 如果不存在就显示map类型的
	fmt.Println(m1["ooo"]) //0
	fmt.Println("---------------------------------------------------")

	// map的遍历
	for k, v := range m1 {
		println(k, v)
	}
	fmt.Println("---------------------------------------------------")

	//map的遍历 遍历key
	for k := range m1 {
		println(k)
	}
	fmt.Println("---------------------------------------------------")

	//map的遍历 遍历value
	for _, v := range m1 {
		println(v)
	}
	fmt.Println("---------------------------------------------------")

	// 删除map的元素  如果删除不存在的  不会报错
	delete(m1, "1231")

	fmt.Println(m1)

}
