package queue

import (
	"context"
	"fmt"
	"sync"

	"github.com/varlyapp/varlyapp/backend/fs"
)

type Pool struct {
	// TaskCount      int
	WorkerCount    int
	TasksChannel   chan Task
	ResultsChannel chan Result
	Done           chan struct{}
}

type TaskID string
type TaskType string
type TaskCallback func(ctx context.Context, t Task) (interface{}, error)
type TaskMetadata struct {
	Dir    string
	Order  []string
	Layers map[string][]fs.Layer
	Width  int
	Height int
	Size   int
}

type Task struct {
	ID        TaskID
	Type      TaskType
	Metadata  TaskMetadata
	Callback  TaskCallback
	Arguments interface{}
}

type Result struct {
	Task  Task
	Value interface{}
	Err   error
}

func (t Task) execute(ctx context.Context) Result {
	value, err := t.Callback(ctx, t)
	if err != nil {
		return Result{
			Task: t,
			Err:  err,
		}
	}

	return Result{
		Task:  t,
		Value: value,
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, tasks <-chan Task, results chan<- Result) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				return
			}
			// fan-in task execution multiplexing results into the results channel
			results <- task.execute(ctx)
		case <-ctx.Done():
			fmt.Printf("cancelled worker. Error detail: %v\n", ctx.Err())
			results <- Result{
				Err: ctx.Err(),
			}
			return
		}
	}
}

func New(workerCount int) Pool {
	return Pool{
		// TaskCount:      taskCount,
		WorkerCount:    workerCount,
		TasksChannel:   make(chan Task, workerCount),
		ResultsChannel: make(chan Result, workerCount),
		Done:           make(chan struct{}),
	}
}

func (pool Pool) Run(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < pool.WorkerCount; i++ {
		wg.Add(1)
		// fan out worker goroutines
		//reading from pools channel and
		//pushing calcs into results channel
		go worker(ctx, &wg, pool.TasksChannel, pool.ResultsChannel)
	}

	wg.Wait()
	close(pool.Done)
	close(pool.ResultsChannel)
}

func (pool Pool) Results() <-chan Result {
	return pool.ResultsChannel
}

func (pool Pool) GenerateFrom(tasks []Task) {
	defer close(pool.TasksChannel)
	for i, _ := range tasks {
		pool.TasksChannel <- tasks[i]
	}
}
