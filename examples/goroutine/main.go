package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Hello, playground")
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup // 计数信号量，记录并维护运行的goroutine
	wg.Add(2)

	// 基于调度器的内部算法，一个正运行的goroutine在工作结束前可以被被停止并重新调度， 以防止某个goroutine长时间占用逻辑处理器CPU
	fmt.Println("Start Goroutines")
	go func() {
		defer wg.Done() // 使用defer声明在函数退出时调用Done方法， defer会修改函数调用时机，在正在执行的函数返回时才真正调用defer声明的函数
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Println("%c a", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Println("%c A ", char)
			}
		}
	}()
	fmt.Println("Waiting To Finish")
	wg.Wait()
}