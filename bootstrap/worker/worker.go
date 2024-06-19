package worker

import (
	"sync"
)

type Task func()

type WorkerPool struct {
	maxWorker    int
	tasksChannel chan Task
	wg           sync.WaitGroup
}

func NewWorkerPool(maxWorker int) WorkerPool {
	return WorkerPool{
		maxWorker:    maxWorker,
		tasksChannel: make(chan Task),
	}
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		go func(workerID int) {
			for task := range wp.tasksChannel {
				task()
				wp.wg.Done()
			}
		}(i + 1)
	}
}

func (wp *WorkerPool) Enqueue(task Task) {
	wp.wg.Add(1)
	wp.tasksChannel <- task
}

func (wp *WorkerPool) Close() {
	wp.wg.Wait()
	close(wp.tasksChannel)
}
