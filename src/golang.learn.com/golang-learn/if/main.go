package main

import "fmt"

// if 表达式1 {
// 	分支1
// } else if 表达式2 {
// 	分支2
// } else{
// 	分支3
// }

// for 初始语句;条件表达式;结束语句{
// 	循环体语句
// }

// func switchDemo1() {
// 	finger := 3
// 	switch finger {
// 	case 1:
// 		fmt.Println("大拇指")
// 	case 2:
// 		fmt.Println("食指")
// 	case 3:
// 		fmt.Println("中指")
// 	case 4:
// 		fmt.Println("无名指")
// 	case 5:
// 		fmt.Println("小拇指")
// 	default:
// 		fmt.Println("无效的输入！")
// 	}
// }

func main() {
	d1 := 102
	if d1 <= 100 {
		fmt.Printf("select number is: %d\n", d1)
	} else if d1 > 100 {
		fmt.Print("large number!")
	}
}
