package main

import (
	"encoding/json"
	"fmt"
)

// 结构体
// 结构体占用一块连续的内存。
// 必须初始化结构体的所有字段。
// 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
// 该方式不能和键值初始化方式混用。
// 使用type和struct关键字来定义结构体，具体代码格式如下：

// type 类型名 struct {
//     字段名 字段类型
//     字段名 字段类型
//     …
// }
// 其中：

// 类型名：标识自定义结构体的名称，在同一个包内不能重复。
// 字段名：表示结构体字段名。结构体中的字段名必须唯一。
// 字段类型：表示结构体字段的具体类型

// 只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。

// 结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。

// var 结构体实例 结构体类型

// 在定义一些临时数据结构等场景下还可以使用匿名结构体。
// var user struct{Name string; Age int}
//     user.Name = "小王子"
//     user.Age = 18

// 我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。 格式如下：

// var p2 = new(person)
// fmt.Printf("%T\n", p2)     //*main.person
// fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}

// 使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。

// p5 := person{
// 	name: "小王子",
// 	city: "北京",
// 	age:  18,
// }
// fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}
// 也可以对结构体指针进行键值对初始化，例如：

// p6 := &person{
// 	name: "小王子",
// 	city: "北京",
// 	age:  18,
// }
// fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}
// 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。

// 方法和接收者
// Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。

// 方法的定义格式如下：

// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//     函数体
// }
// 其中，

// 接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
// 接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
// 方法名、参数列表、返回参数：具体格式与函数定义相同。

type Address struct {
	City   string `json:"city"`
	Detail string `json:"detail"`
}

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender bool   `json:"gender"`
	Type   bool   `json:"type"`
	Address
}

func main() {
	p1 := Person{
		Name:   "gg",
		Age:    12,
		Gender: false,
		Type:   true,
	}
	p1.Name = "ss"
	fmt.Printf("this person is : %#v\n", p1)
	p2 := newPerson("aa", 15, false, false)
	fmt.Printf("this person is : %#v\n", p2)
	p2.Name = "cc"
	fmt.Printf("this person is : %#v\n", p2)
	p2.Speak()
	data, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("josn data is: %s\n", data)
	p3 := &Person{}
	err = json.Unmarshal([]byte(data), p3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("json parse person is : %#v\n", p3)
}

func newPerson(name string, age int, gender, _type bool) *Person {
	return &Person{
		Name:   name,
		Age:    age,
		Gender: gender,
		Type:   _type,
	}
}

func (p *Person) Speak() {
	fmt.Printf("my name is: %s\n", p.Name)
}
