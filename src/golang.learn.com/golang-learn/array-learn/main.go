package main

import "fmt"

// var 数组变量名 [元素数量]T
// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
// 注意：
// 数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
// [n]*T表示指针数组，*[n]T表示数组指针 。
func main() {
	a := [...]int{1, 2, 3, 4, 5}
	// a1 := [...]string{"北京", "上海", "深圳"}
	// fmt.Println(a1)
	// 方法1：for循环遍历
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// 方法2：for range遍历
	// for index, value := range a {
	// 	fmt.Println(index, value)
	// }

	// a2 := [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	// fmt.Println(a2) //[[北京 上海] [广州 深圳] [成都 重庆]]
	// fmt.Println(a2[2][1]) //支持索引取值:重庆
	// for i:= 0;i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	for _,v := range a {
		fmt.Println(v)
	}
	a2 := [...]bool{false, true, false, false, true}
	fmt.Printf("length of ary is : %d\n content of ary  is : %v\n", len(a2), a2)
	a3 := a[1:3:5] // 切片
	fmt.Print(a3)
}
