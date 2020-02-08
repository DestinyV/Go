package main

import "fmt"

// var name []T
// name:表示变量名
// T:表示切片中的元素类型

// 切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

// make([]T, size, cap)
// T:切片的元素类型
// size:切片中元素的数量
// cap:切片的容量

// append()方法为切片添加元素
// append()函数将元素追加到切片的最后并返回该切片。
// 切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。

// 使用copy()函数复制切片 copy(destSlice, srcSlice []T)
// srcSlice: 数据来源切片
// destSlice: 目标切片

// 从切片中删除元素
// 要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
func main() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	// var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil
	d := make([]int, 1, 10)
	d[0] = 10
	d = append(d, 1, 2, 3, 4)
	e := d[1:]
	f := make([]int, 5)
	copy(f, e)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
