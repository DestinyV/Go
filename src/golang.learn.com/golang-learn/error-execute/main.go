package main

import (
	"fmt"
)

func main() {
err()
 f1()
}

func f1()  {
	fmt.Println("execute f1");
}

func err()  {
	defer func(){
		err := recover()
		fmt.Println(err) // recover 后面程序还会继续执行
		fmt.Println("release cache...");
	}()
 panic("error!")
}
