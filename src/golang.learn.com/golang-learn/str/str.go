package main

import "fmt"
import "strings"

func main() {
	s1 := "h-e-l-l-o"
	s2 := "world"
	fmt.Printf("string is: %s\ntype is: %T\n", s1, s1)
	byte1 := []byte(s2)
	rune1 := []rune(s2)
	for i := 0; i < len(s2); i++ {
		fmt.Printf("current item is %d\n", byte1[i])
	}
	for _, v := range rune1 {
		fmt.Printf("current item is %c\n", v)
	}
	// 分隔
	fmt.Println(strings.Split(s1, "-"))
}
