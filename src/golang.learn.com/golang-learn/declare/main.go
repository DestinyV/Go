package main

import "fmt"

// 全局变量可以不用使用
var (
	age    int
	name   string
	gender bool
	area   float32
)

// 全局变量可以不用使用
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	level := KB // 局部变量申明后必须被使用
	fmt.Printf("level is :%d\n", level)
}
