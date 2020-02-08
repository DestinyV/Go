package main

import (
	"fmt"
	"reflect"
)

// 反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。

// 支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

// Go程序在运行期使用reflect包访问程序的反射信息。

// reflect包
// 在Go语言的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组成的(我们在上一篇接口的博客中有介绍相关概念)。 在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。

// type Kind uint
// const (
//     Invalid Kind = iota  // 非法类型
//     Bool                 // 布尔型
//     Int                  // 有符号整型
//     Int8                 // 有符号8位整型
//     Int16                // 有符号16位整型
//     Int32                // 有符号32位整型
//     Int64                // 有符号64位整型
//     Uint                 // 无符号整型
//     Uint8                // 无符号8位整型
//     Uint16               // 无符号16位整型
//     Uint32               // 无符号32位整型
//     Uint64               // 无符号64位整型
//     Uintptr              // 指针
//     Float32              // 单精度浮点数
//     Float64              // 双精度浮点数
//     Complex64            // 64位复数类型
//     Complex128           // 128位复数类型
//     Array                // 数组
//     Chan                 // 通道
//     Func                 // 函数
//     Interface            // 接口
//     Map                  // 映射
//     Ptr                  // 指针
//     Slice                // 切片
//     String               // 字符串
//     Struct               // 结构体
//     UnsafePointer        // 底层指针
// )

// ValueOf
// reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。

// reflect.Value类型提供的获取原始值的方法如下：

// 方法	说明
// Interface() interface {}	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
// Int() int64	将值以 int 类型返回，所有有符号整型均可以此方式返回
// Uint() uint64	将值以 uint 类型返回，所有无符号整型均可以此方式返回
// Float() float64	将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
// Bool() bool	将值以 bool 类型返回
// Bytes() []bytes	将值以字节数组 []bytes 类型返回
// String() string	将值以字符串类型返回

func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}
