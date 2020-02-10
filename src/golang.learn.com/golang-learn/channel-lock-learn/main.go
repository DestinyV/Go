package main

// 并发安全和锁
// 有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）。类比现实生活中的例子有十字路口被各个方向的的汽车竞争；还有火车上的卫生间被车厢里的人竞争。

// 举个例子：

// var x int64
// var wg sync.WaitGroup

// func add() {
// 	for i := 0; i < 5000; i++ {
// 		x = x + 1
// 	}
// 	wg.Done()
// }
// func main() {
// 	wg.Add(2)
// 	go add()
// 	go add()
// 	wg.Wait()
// 	fmt.Println(x)
// }
// 上面的代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符。

// 互斥锁
// 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁。 使用互斥锁来修复上面代码的问题：

// var x int64
// var wg sync.WaitGroup
// var lock sync.Mutex

// func add() {
// 	for i := 0; i < 5000; i++ {
// 		lock.Lock() // 加锁
// 		x = x + 1
// 		lock.Unlock() // 解锁
// 	}
// 	wg.Done()
// }
// func main() {
// 	wg.Add(2)
// 	go add()
// 	go add()
// 	wg.Wait()
// 	fmt.Println(x)
// }
// 使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。

// 读写互斥锁
// 互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。

// 读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

// 读写锁示例：

// var (
// 	x      int64
// 	wg     sync.WaitGroup
// 	lock   sync.Mutex
// 	rwlock sync.RWMutex
// )

// func write() {
// 	// lock.Lock()   // 加互斥锁
// 	rwlock.Lock() // 加写锁
// 	x = x + 1
// 	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
// 	rwlock.Unlock()                   // 解写锁
// 	// lock.Unlock()                     // 解互斥锁
// 	wg.Done()
// }

// func read() {
// 	// lock.Lock()                  // 加互斥锁
// 	rwlock.RLock()               // 加读锁
// 	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
// 	rwlock.RUnlock()             // 解读锁
// 	// lock.Unlock()                // 解互斥锁
// 	wg.Done()
// }

// func main() {
// 	start := time.Now()
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go write()
// 	}

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go read()
// 	}

// 	wg.Wait()
// 	end := time.Now()
// 	fmt.Println(end.Sub(start))
// }
// 需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。

// sync.WaitGroup
// 在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。 sync.WaitGroup有以下几个方法：

// 方法名	功能
// (wg * WaitGroup) Add(delta int)	计数器+delta
// (wg *WaitGroup) Done()	计数器-1
// (wg *WaitGroup) Wait()	阻塞直到计数器变为0
// sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。

// 我们利用sync.WaitGroup将上面的代码优化一下：

// var wg sync.WaitGroup

// func hello() {
// 	defer wg.Done()
// 	fmt.Println("Hello Goroutine!")
// }
// func main() {
// 	wg.Add(1)
// 	go hello() // 启动另外一个goroutine去执行hello函数
// 	fmt.Println("main goroutine done!")
// 	wg.Wait()
// }
// 需要注意sync.WaitGroup是一个结构体，传递的时候要传递指针。

// sync.Once
// 说在前面的话：这是一个进阶知识点。

// 在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。

// Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。

// sync.Once只有一个Do方法，其签名如下：

// func (o *Once) Do(f func()) {}
// 备注：如果要执行的函数f需要传递参数就需要搭配闭包来使用。

// 加载配置文件示例
// 延迟一个开销很大的初始化操作到真正用到它的时候再执行是一个很好的实践。因为预先初始化一个变量（比如在init函数中完成初始化）会增加程序的启动耗时，而且有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做的。我们来看一个例子：

// var icons map[string]image.Image

// func loadIcons() {
// 	icons = map[string]image.Image{
// 		"left":  loadIcon("left.png"),
// 		"up":    loadIcon("up.png"),
// 		"right": loadIcon("right.png"),
// 		"down":  loadIcon("down.png"),
// 	}
// }

// // Icon 被多个goroutine调用时不是并发安全的
// func Icon(name string) image.Image {
// 	if icons == nil {
// 		loadIcons()
// 	}
// 	return icons[name]
// }
// 多个goroutine并发调用Icon函数时不是并发安全的，现代的编译器和CPU可能会在保证每个goroutine都满足串行一致的基础上自由地重排访问内存的顺序。loadIcons函数可能会被重排为以下结果：

// func loadIcons() {
// 	icons = make(map[string]image.Image)
// 	icons["left"] = loadIcon("left.png")
// 	icons["up"] = loadIcon("up.png")
// 	icons["right"] = loadIcon("right.png")
// 	icons["down"] = loadIcon("down.png")
// }
// 在这种情况下就会出现即使判断了icons不是nil也不意味着变量初始化完成了。考虑到这种情况，我们能想到的办法就是添加互斥锁，保证初始化icons的时候不会被其他的goroutine操作，但是这样做又会引发性能问题。

// 使用sync.Once改造的示例代码如下：

// var icons map[string]image.Image

// var loadIconsOnce sync.Once

// func loadIcons() {
// 	icons = map[string]image.Image{
// 		"left":  loadIcon("left.png"),
// 		"up":    loadIcon("up.png"),
// 		"right": loadIcon("right.png"),
// 		"down":  loadIcon("down.png"),
// 	}
// }

// // Icon 是并发安全的
// func Icon(name string) image.Image {
// 	loadIconsOnce.Do(loadIcons)
// 	return icons[name]
// }
// 并发安全的单例模式
// 下面是借助sync.Once实现的并发安全的单例模式：

// package singleton

// import (
//     "sync"
// )

// type singleton struct {}

// var instance *singleton
// var once sync.Once

// func GetInstance() *singleton {
//     once.Do(func() {
//         instance = &singleton{}
//     })
//     return instance
// }
// sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。

// sync.Map
// Go语言中内置的map不是并发安全的。请看下面的示例：

// var m = make(map[string]int)

// func get(key string) int {
// 	return m[key]
// }

// func set(key string, value int) {
// 	m[key] = value
// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			set(key, n)
// 			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }
// 上面的代码开启少量几个goroutine的时候可能没什么问题，当并发多了之后执行上面的代码就会报fatal error: concurrent map writes错误。

// 像这种场景下就需要为map加锁来保证并发的安全性了，Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。

// var m = sync.Map{}

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			m.Store(key, n)
// 			value, _ := m.Load(key)
// 			fmt.Printf("k=:%v,v:=%v\n", key, value)
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }
// 原子操作
// 代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好。Go语言中原子操作由内置的标准库sync/atomic提供。

// atomic包
// 方法	解释
// func LoadInt32(addr *int32) (val int32)
// func LoadInt64(addr *int64) (val int64)
// func LoadUint32(addr *uint32) (val uint32)
// func LoadUint64(addr *uint64) (val uint64)
// func LoadUintptr(addr *uintptr) (val uintptr)
// func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)	读取操作
// func StoreInt32(addr *int32, val int32)
// func StoreInt64(addr *int64, val int64)
// func StoreUint32(addr *uint32, val uint32)
// func StoreUint64(addr *uint64, val uint64)
// func StoreUintptr(addr *uintptr, val uintptr)
// func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)	写入操作
// func AddInt32(addr *int32, delta int32) (new int32)
// func AddInt64(addr *int64, delta int64) (new int64)
// func AddUint32(addr *uint32, delta uint32) (new uint32)
// func AddUint64(addr *uint64, delta uint64) (new uint64)
// func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)	修改操作
// func SwapInt32(addr *int32, new int32) (old int32)
// func SwapInt64(addr *int64, new int64) (old int64)
// func SwapUint32(addr *uint32, new uint32) (old uint32)
// func SwapUint64(addr *uint64, new uint64) (old uint64)
// func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
// func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)	交换操作
// func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
// func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
// func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
// func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
// func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
// func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)	比较并交换操作
示例
我们填写一个示例来比较下互斥锁和原子操作的性能。

type Counter interface {
	Inc()
	Load() int64
}

// 普通版
// type CommonCounter struct {
// 	counter int64
// }

// func (c CommonCounter) Inc() {
// 	c.counter++
// }

// func (c CommonCounter) Load() int64 {
// 	return c.counter
// }

// 互斥锁版
// type MutexCounter struct {
// 	counter int64
// 	lock    sync.Mutex
// }

// func (m *MutexCounter) Inc() {
// 	m.lock.Lock()
// 	defer m.lock.Unlock()
// 	m.counter++
// }

// func (m *MutexCounter) Load() int64 {
// 	m.lock.Lock()
// 	defer m.lock.Unlock()
// 	return m.counter
// }

// 原子操作版
// type AtomicCounter struct {
// 	counter int64
// }

// func (a *AtomicCounter) Inc() {
// 	atomic.AddInt64(&a.counter, 1)
// }

// func (a *AtomicCounter) Load() int64 {
// 	return atomic.LoadInt64(&a.counter)
// }

// func test(c Counter) {
// 	var wg sync.WaitGroup
// 	start := time.Now()
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			c.Inc()
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	end := time.Now()
// 	fmt.Println(c.Load(), end.Sub(start))
// }

// func main() {
// 	c1 := CommonCounter{} // 非并发安全
// 	test(c1)
// 	c2 := MutexCounter{} // 使用互斥锁实现并发安全
// 	test(&c2)
// 	c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
// 	test(&c3)
// }
// atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。

func main()  {
	
}

