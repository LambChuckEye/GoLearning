package main

import (
	"fmt"
	"sync"
	"math"
)

func main() {
	//waitgroup用于协程并发控制
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		//标记任务开始
		wg.Add(1)
		go func(a int) {
			fmt.Println(fmt.Sprintf("work %d exec", a))
			//标记任务结束
			wg.Done()
		}(i)
	}
	//主线程等待任务结束
	wg.Wait()
	fmt.Println("main end")
	math.MaxInt
	fmt.Println(math.MaxInt)
}
