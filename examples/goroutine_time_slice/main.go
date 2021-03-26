package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup // 计数信号量，记录并维护运行的goroutine
func main() {
	fmt.Println("Hello, playground")
	//给每个可用的和兴分配一个逻辑处理器
	runtime.GOMAXPROCS(runtime.NumCPU)


	wg.Add(2)

	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")
	fmt.Println("Waiting To Finish")
	wg.Wait()

}

func printPrime(prefix string) {
	defer wg.Done()
next:
	for count := 2; count < 5000; count++ {
		for char := 2; char < count; char++ {
			if count%char == 0 {
				continue next
			}
			fmt.Println("%s:%d\n", prefix, count)
		}

	}
	fmt.Println("Completed", prefix)
}