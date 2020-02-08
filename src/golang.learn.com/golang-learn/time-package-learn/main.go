package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now time is", now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 时间戳
	timestamp := now.Unix()
	nanostamp := now.UnixNano()
	fmt.Println("timestamp is ", timestamp)
	fmt.Println("nanostamp is ", nanostamp)

	// 转换为时间格式
	t1 := time.Unix(1581158424, 0)
	fmt.Println(t1)

	// 时间间隔
	oneSec := time.Second
	fmt.Println("oneSec", oneSec)

	// 时间操作
	// func (t Time) Add(d Duration) Time
	fmt.Println("add 24hours", now.Add(time.Hour*24))
	// func (t Time) Sub(u Time) Duration
	// func (t Time) Equal(u Time) bool
	// func (t Time) Before(u Time) bool
	// func (t Time) After(u Time) bool

	// time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道（channel）。
	// ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	// for i := range ticker {
	// 	fmt.Println(i) //每秒都会执行的任务
	// }

	// 时间类型有一个自带的方法Format进行格式化，需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。也许这就是技术人员的浪漫吧。
	// 补充：如果想格式化为12小时方式，需指定PM。
	// now := time.Now()
	// // 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// // 24小时制
	// fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// // 12小时制
	// fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	// fmt.Println(now.Format("2006/01/02 15:04"))
	// fmt.Println(now.Format("15:04 2006/01/02"))
	// fmt.Println(now.Format("2006/01/02"))

	// 解析字符串格式的时间
	// 加载时区
	// loc, err := time.LoadLocation("Asia/Shanghai")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // 按照指定时区和指定格式解析字符串时间
	// timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//sleep
	fmt.Println("before sleep")
	time.Sleep(time.Second)
	fmt.Println("after sleep")

}
