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
	court := make(chan int)

	// 计数 + 2 , 表示要等待两个goroutine
	wg.Add(2)
	// 启动两个选手
	go player("first1", court)
	go player("two2", court)
	//发球
	court <- 1

	//等待游戏结束
	wg.Wait()
}

// player
func player(name string, court chan int) {
	defer wg.Done()
	for {
		//等待球被击打过来
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won \n", name)
			return
		}
		// 选随机数， 用于判断是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed \n", name)
			// 关闭通道， 表示输了
			close(court)
			return
		}
		//显示击球数，将击球数+1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		// 将球打向对手
		court <- ball
	}
}
