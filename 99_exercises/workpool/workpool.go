package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Task struct {
	ID      int
	Payload string
	Status  string
	Error   string
}

func ProcessTask(task *Task) error {
	time.Sleep(time.Second)
	if task.ID%7 == 0 {
		return fmt.Errorf("Task %d could not be processed", task.ID)
	}
	return nil
}

type Context struct {
	Cancelled bool
	Timeout   int
}

type Worker struct {
	ID        int
	RequestCh chan *Task
	ResultCh  chan *Task
	Running   bool
}

func (w *Worker) Run(ctx *Context) {
	w.Running = true
	for {
		select {
		case task := <-w.RequestCh:
			fmt.Printf("Worker %d processing task...\n", w.ID)
			err := ProcessTask(task)
			if err != nil {
				task.Status = "failed"
				task.Error = err.Error()
			} else {
				task.Status = "success"
			}
			if !ctx.Cancelled {
				w.ResultCh <- task
			} else {
				fmt.Println("Cancellation Requested for worker:", w.ID)
				w.Running = false
				return
			}
			fmt.Printf("Worker %d done\n", w.ID)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	workerCount := 10
	workerPool := []*Worker{}

	requestCh := make(chan *Task, 100)
	resultCh := make(chan *Task)

	ctx := &Context{
		Cancelled: false,
	}

	for i := range workerCount {
		worker := &Worker{
			ID:        i,
			RequestCh: requestCh,
			ResultCh:  resultCh,
		}
		workerPool = append(workerPool, worker)
		go worker.Run(ctx)
	}

	for range 100 {
		task := &Task{
			ID:      rand.IntN(100),
			Payload: "Hello World",
		}
		fmt.Println("Sending task:", task.ID)
		requestCh <- task
	}

	for i := 0; i < 50; i++ {
		select {
		case task := <-resultCh:
			fmt.Printf(
				"Task %d processed. Status: %s. Error: %s.\n", task.ID, task.Status, task.Error,
			)
		}
	}

	ctx.Cancelled = true
	time.Sleep(5 * time.Second)

	close(requestCh)
	close(resultCh)

	time.Sleep(2 * time.Second)
	for _, worker := range workerPool {
		if worker.Running == true {
			panic("Worker shouldn't be running")
		}
	}
}
