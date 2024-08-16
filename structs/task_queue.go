package structs

import "sync"

type taskQueue struct {
	enqueue   chan Task
	dequeue   chan Task
	waitGroup sync.WaitGroup
}

// NewTaskQueue creates a new taskQueue struct.
func NewTaskQueue() *taskQueue {
	return &taskQueue{
		enqueue: make(chan Task),
		dequeue: make(chan Task),
	}
}

// Start continuously sends tasks from the enqueue channel into the dequeue channel
func (tQ *taskQueue) Start() {
	go func() {
		// Send tasks from enqueue channel into the dequeue channel
		for task := range tQ.enqueue {
			// Some features for verifying the task can wrap around the following
			tQ.dequeue <- task
		}
	}()
}

// SubmitTask submits a task to the task queue.
func (tQ *taskQueue) SubmitTask(t Task) {
	tQ.enqueue <- t
}

// Wait calls the Wait() function for the task queue's wait group.
func (tQ *taskQueue) Wait() {
	tQ.waitGroup.Wait()
}
