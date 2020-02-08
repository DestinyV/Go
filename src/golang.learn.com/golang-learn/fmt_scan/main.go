package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmtScan()
	// scanByBufio()
	// 写出
	// fmt.Fprintln(os.Stdout, "this is a log information")
	fileObj, _ := os.OpenFile("./log.txt", os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
}

func fmtScan() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}

func scanByBufio() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("what you input is : %s\n", string(input))
}
