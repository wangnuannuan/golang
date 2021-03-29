package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 用来等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	//创建无缓冲通道
	baton := make(chan int)

	// 为最后一位跑步者计数+1
	wg.Add(1)

	//第一位跑步者持有接力棒
	go Runner(baton)

	// 开始比赛
	baton <- 1

	//等待比赛结束
	wg.Wait()
}

// Runner 模拟接力比赛中的第一位跑步者

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton
	// 开始绕着跑道跑
	fmt.Printf("Runner %d Running with Baton \n", runner)
	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the Line\n", newRunner)
	}
	time.Sleep(100*time.Millisecond)
	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}
	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange with Runner %d\n", runner, newRunner)
	baton <- newRunner
}

