package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Go语言中 map的定义语法如下：

// map[KeyType]ValueType
// 其中，

// KeyType:表示键的类型。
// ValueType:表示键对应的值的类型。
// map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

// make(map[KeyType]ValueType, [cap])
// Go语言中有个判断map中键是否存在的特殊写法，格式如下:
// value, ok := map[key]
// 使用delete()函数删除键值对
// 使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：

// delete(map, key)
// 其中，
// map:表示要删除键值对的map
// key:表示要删除的键值对的键
func main() {
	// m1 := map[string][]int{}
	// m1["first"] = make([]int, 0, 10)
	// m1["first"] = append(m1["first"], 125)
	// fmt.Printf("type is : %T\ncontent is : %v\n", m1, m1)
	// _, exist := m1["first"]
	// if exist {
	// 	fmt.Println(m1["first"])
	// } else {
	// 	fmt.Println("not exist")
	// }

	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	m2 := make(map[string]int, 50)   // Map
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("edu%02d", i)
		value := rand.Intn(100)
		m2[key] = value
	}
	fmt.Println(m2)
	s1 := make([]string, 0, 50) // 切片存放所有key值
	for key := range m2 {
		s1 = append(s1, key)
	}
	sort.Strings(s1) // 给切片内的值排序
	for _, key := range s1 {
		fmt.Println(key, m2[key])
	}
}
