package lib

import (
	"context"
	"runtime"
	"sync"
)

type Job struct {
	Id     int
	Config interface{}
}

type Callback func(ctx context.Context, id int, job Job)

func Batch(ctx context.Context, size int, jobs []Job, fn Callback) {
	var (
		wg              sync.WaitGroup
		jobChannel      = make(chan Job)
		numberOfWorkers = size
	)

	// if number of workers is 0, use a sensible default
	if numberOfWorkers == 0 {
		numberOfWorkers = runtime.NumCPU() * 3
	}

	wg.Add(numberOfWorkers)

	// start the workers
	for i := 0; i < numberOfWorkers; i++ {
		go work(ctx, i, &wg, jobChannel, fn)
	}

	// send jobs to worker
	for _, job := range jobs {
		jobChannel <- job
	}

	close(jobChannel)
	wg.Wait()
}

func work(ctx context.Context, id int, wg *sync.WaitGroup, jobChannel <-chan Job, fn Callback) {
	defer wg.Done()
	for job := range jobChannel {
		fn(ctx, id, job)
	}
}
