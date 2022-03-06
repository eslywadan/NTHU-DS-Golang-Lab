package workerpool

import (
		"context"
		"sync"
)

type Task struct {
		Func func(args ...interface{}) *Result
		Args []interface{}
}

type Result struct {
	Value interface{}
	Err error
}

type WorkerPool interface {
	Start(ctx context.Context)
	Tasks() chan *Task
	Results() chan *Result
}

type workerPool struct {
	numWorkers int
	tasks 	 chan *Task
	results  chan *Result
	wg 		 *sync.WaitGroup
}

var _ WorkerPool = (*workerPool)(nil)

func NewWorkerPool(numWorkers int, bufferSize int) *workerPool {
	return &workerPool{
			numWorkers: numWorkers,
			tasks:	make(chan *Task,  bufferSize),
			results: make(chan *Result, bufferSize),
			wg: 	&sync.WaitGroup{},
	}
}

func (wp *workerPool) Start(ctx context.Context) {
		// TODO: implementation
		//
		// Starts numWorker of goroutines, wait untill all jobs are done.
		// Remember to close the result channel before exit.
		
}

func (wp *workerPool) Tasks() chan *Task {
		return wp.tasks
}

func (wp *workerPool) Results() chan *Result {
		return wp.results
}

func (wp *workerPool) run(ctx context.Context) {
		// TODO: Implementation
		//
		// Keeps fetching task from the task channel, do the task.
		// Then makes sure to exit if context is done.
}