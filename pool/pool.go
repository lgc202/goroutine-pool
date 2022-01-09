package pool

import "fmt"

type Pool struct {
	// Task队列
	JobsChannel chan *Task

	// worker的数量
	workerNum int
}

func NewPool(cap int) *Pool {
	return &Pool{
		JobsChannel: make(chan *Task),
		workerNum:   cap,
	}
}

// 协程池创建一个worker，并且让这个worker去工作
func (p *Pool) worker(workerID int) {
	// 永久从JobChannel里面取任务
	for {
		select {
		case task := <-p.JobsChannel:
			task.Execute()
			fmt.Println("worker ID ", workerID, " 执行完了一个任务")
		}
	}
}

// Run 让协程池开始工作
func (p *Pool) Run() {
	// 根据workerNum创建worker
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}
}
