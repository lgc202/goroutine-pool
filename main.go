package main

import (
	"fmt"
	"goroutine-pool/pool"
	"time"
)

func main() {
	// 1. 创建一些任务
	task := pool.NewTask(func() {
		fmt.Println(time.Now())
		return
	})

	// 2. 创建协程池
	p := pool.NewPool(4)

	// 3. 将这些任务交给协程池
	go func() {
		for {
			p.JobsChannel <- task
		}
	}()

	// 4. 启动pool
	p.Run()
	select {}
}
