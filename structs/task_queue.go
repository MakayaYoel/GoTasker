package structs

import "sync"

type taskQueue struct {
	tasks     []Task
	enqueue   chan Task
	dequeue   chan Task
	waitGroup sync.WaitGroup
}

// NewTaskQueue creates a new taskQueue struct.
func NewTaskQueue() *taskQueue {
	return &taskQueue{
		tasks:   make([]Task, 0),
		enqueue: make(chan Task),
		dequeue: make(chan Task),
	}
}

// Start continuously appends "tasks" from the enqueue channel into the tasks slice, or sends tasks from the "tasks" slice into the dequeue channel.
func (tQ *taskQueue) Start() {
	go func() {
		for {
			select {
			// Send tasks from enqueue channel into the "tasks" slice
			case task := <-tQ.enqueue:
				tQ.tasks = append(tQ.tasks, task)
			default:
				// Send tasks from the "tasks" slice into the dequeue channel
				if len(tQ.tasks) > 0 {
					tQ.dequeue <- tQ.tasks[0]
					tQ.tasks = tQ.tasks[1:]
				}
			}
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
