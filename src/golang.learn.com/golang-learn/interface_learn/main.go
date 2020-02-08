package main

import "fmt"

// Go语言提倡面向接口编程。

// 每个接口由数个方法组成，接口的定义格式如下：

// type 接口类型名 interface{
//     方法名1( 参数列表1 ) 返回值列表1
//     方法名2( 参数列表2 ) 返回值列表2
//     …
// }
// 其中：

// 接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
// 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
// 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

// type writer interface {
// 	Write([]byte) error
// }

// 空接口作为map的值
// 使用空接口实现可以保存任意值的字典。

// // 空接口作为map值
// 	var studentInfo = make(map[string]interface{})
// 	studentInfo["name"] = "沙河娜扎"
// 	studentInfo["age"] = 18
// 	studentInfo["married"] = false

type voice interface {
	Translate(string) string
	Talk(string)
	Speak(string)
}

type mammal struct {
	name    string
	species string
}

func (m *mammal) Translate(a string) string {
	fmt.Println(a)
	return a + "[Translated]"
}

func main() {
	m1 := &mammal{
		name:    "aa",
		species: "human",
	}
	s1 := m1.Translate("hello")
	fmt.Println(s1)
}

// 想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：

// x.(T)
// 其中：

// x：表示类型为interface{}的变量
// T：表示断言x可能是的类型。
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
