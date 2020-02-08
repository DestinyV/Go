package main 

import(
	"fmt"
	"reflect"
)

func main() {
	var (
		t1 int8 = 12 // byte
		t2 float32 = 12.25
	)
	fmt.Printf("type of t1 is: %s\n", reflect.TypeOf(t1))
	fmt.Printf("type of t2 is: %s\n", reflect.TypeOf(t2))
}