package main

import "fmt"

// new是一个内置的函数，它的函数签名如下：
// func new(Type) *Type
// Type表示类型，new函数只接受一个参数，这个参数是一个类型
// *Type表示类型指针，new函数返回一个指向该类型内存地址的指针。

// new与make的区别
// 二者都是用来做内存分配的。
// make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
// 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

func main() {
	a := 13
	addr := &a
	fmt.Printf("type is : %T\naddress is : %v\npointer is: %p\nvalue is : %d\n", addr, addr, addr, *addr)
	var b int
	f1(&b)
	fmt.Printf("f1 return is:%d\n", b)
}

func f1(p *int) {
	*p = 123
}
