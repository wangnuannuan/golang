package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go build -race 用竞争检测器标志来编译程序
var (
	// counter 时所有goroutine都要增加其值的变量
	counter int
	// 等待程序结束
	wg sync.WaitGroup
)

func main() {
	// var wg sync.WaitGroup // 计数信号量，记录并维护运行的goroutine
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)

	//等待goroutine结束
	wg.Wait()
	fmt.Println("Final Counter:", counter) // 2
}

//incCounter增加包里counter变量的值

func incCounter(id int) {
	defer wg.Done() // 使用defer声明在函数退出时调用Done方法， defer会修改函数调用时机，在正在执行的函数返回时才真正调用defer声明的函数
	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched() //将goroutine从当前线程退出，给其他goroutine运行的机会。 强制调度器切换两个goroutine,以便让竞争状态的效果更明显
		value++
		counter = value
	}
}