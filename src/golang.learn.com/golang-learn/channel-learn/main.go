package main

import (
	"fmt"
	"time"
)

// channel
// 单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

// 虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

// Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

// 如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

// Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
func main() {
	// 	channel类型
	// channel是一种类型，一种引用类型。声明通道类型的格式如下：

	// var 变量 chan 元素类型
	// 举几个例子：

	// var ch1 chan int   // 声明一个传递整型的通道
	// var ch2 chan bool  // 声明一个传递布尔型的通道
	// var ch3 chan []int // 声明一个传递int切片的通道
	// var ch8 chan string
	// 	创建channel
	// 通道是引用类型，通道类型的空值是nil。

	// var ch chan int
	// fmt.Println(ch) // <nil>

	// 声明的通道后需要使用make函数初始化之后才能使用。

	// 创建channel的格式如下：

	// make(chan 元素类型, [缓冲大小])
	// channel的缓冲大小是可选的。

	// 举几个例子：

	// ch4 := make(chan int)
	// ch5 := make(chan bool)
	// ch6 := make(chan []int)
	// ch9 := make(chan string)

	// 	channel操作
	// 通道有发送（send）、接收(receive）和关闭（close）三种操作。

	// 发送和接收都使用<-符号。

	// 现在我们先使用以下语句定义一个通道：

	// ch7 := make(chan int)
	// 发送
	// 将一个值发送到通道中。

	// ch7 <- 10 // 把10发送到ch中
	// 接收
	// 从一个通道中接收值。

	// x := <-ch7 // 从ch中接收值并赋值给变量x
	// <-ch7      // 从ch中接收值，忽略结果
	// 关闭
	// 我们通过调用内置的close函数来关闭通道。

	// close(ch7)
	// fmt.Println(x)
	// 关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

	// 关闭后的通道有以下特点：

	// 对一个关闭的通道再发送值就会导致panic。
	// 对一个关闭的通道进行接收会一直获取值直到通道为空。
	// 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
	// 关闭一个已经关闭的通道会导致panic。

	ch := make(chan int) // ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。
	// 	无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。

	// 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。
	go receiveChan(ch)
	ch <- 10

	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println("number form ch2:", i)
	}
}

func receiveChan(rec chan int) {
	num := <-rec
	close(rec)
	fmt.Println("receiveChan number:", num)
}

// 单向通道
// 有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。

// Go语言中提供了单向通道来处理这种情况。例如，我们把上面的例子改造如下：

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

// 其中，

// chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
// <-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。
// 在函数传参及任何赋值操作中可以将双向通道转换为单向通道，但反过来是不可以的。

// 通道总结
// channel常见的异常总结，如下图：channel异常总结

// 关闭已经关闭的channel也会引发panic。

// worker pool（goroutine池）
// 在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨。

// 一个简易的work pool示例代码如下：

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

// func main() {
// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)
// 	// 开启3个goroutine
// 	for w := 1; w <= 3; w++ {
// 		go worker(w, jobs, results)
// 	}
// 	// 5个任务
// 	for j := 1; j <= 5; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)
// 	// 输出结果
// 	for a := 1; a <= 5; a++ {
// 		<-results
// 	}
// }
// select多路复用
// 在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。你也许会写出如下代码使用遍历的方式来实现：

// for{
//     // 尝试从ch1接收值
//     data, ok := <-ch1
//     // 尝试从ch2接收值
//     data, ok := <-ch2
//     …
// }
// 这种方式虽然可以实现从多个通道接收值的需求，但是运行性能会差很多。为了应对这种场景，Go内置了select关键字，可以同时响应多个通道的操作。

// // select的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。具体格式如下：

// select{
//     case <-ch1:
//         ...
//     case data := <-ch2:
//         ...
//     case ch3<-data:
//         ...
//     default:
//         默认操作
// }
// 举个小例子来演示下select的使用：

// func main() {
// 	ch := make(chan int, 1)
// 	for i := 0; i < 10; i++ {
// 		select {
// 		case x := <-ch:
// 			fmt.Println(x)
// 		case ch <- i:
// 		}
// 	}
// }
// 使用select语句能提高代码的可读性。

// 可处理一个或多个channel的发送/接收操作。
// 如果多个case同时满足，select会随机选择一个。
// 对于没有case的select{}会一直等待，可用于阻塞main函数。
