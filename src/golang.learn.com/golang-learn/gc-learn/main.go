package main

import (
	"fmt"
	"runtime"
)

func main() {
	getInfor(0) // 15
	getInfor(1) // 10
}

func getInfor(n int) {
	// skip 向外跳入的层数
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("get call information failled")
		return
	}
	// fmt.Println("gc:", pc)
	fmt.Println("file:", file)
	fmt.Println("line:", line)
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println("funcName:", funcName)
}
