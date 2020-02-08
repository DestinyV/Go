package main

import "fmt"

func main() {
	rf1 := f1(2, 3)
	fmt.Println(rf1) // 5
	rf2 := f2(f1)    // 9 defer 10
	rf3 := rf2(6)
	fmt.Println(rf3) // 60
}

func f1(a, b int) int {
	return a + b
}

func f2(add func(int, int) int) func(int) int {
	r1 := add(4, 5)
	rf := func(c int) int {
		fmt.Printf("return function r1: %d\n", r1)
		return c * r1
	}
	defer add1(&r1)
	fmt.Println(r1)
	return rf
}

func add1(a *int) {
	*a++
	fmt.Printf("inner defer: %d\n", *a)
}
