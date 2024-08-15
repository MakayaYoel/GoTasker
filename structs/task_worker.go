package structs

import "log"

type taskWorker struct {
	id        int
	taskQueue *taskQueue
}

// NewTaskWorker creates a new taskWorker struct.
func NewTaskWorker(id int, tQ *taskQueue) *taskWorker {
	return &taskWorker{
		id:        id,
		taskQueue: tQ,
	}
}

// Start continuously takes tasks from its task queue's dequeue channel and executes them.
func (tW *taskWorker) Start() {
	go func() {
		for {
			// Take task from dequeue channel and execute it.
			task := <-tW.taskQueue.dequeue

			tW.taskQueue.waitGroup.Add(1)

			err := task.Execute()

			if err != nil {
				log.Printf("\nTask Worker (ID: %d) has encountered an error while executing a task: %s", tW.id, err.Error())
			} else {
				log.Printf("\nTask Worker (ID: %d) has successfully executed a task.", tW.id)
			}

			tW.taskQueue.waitGroup.Done()
		}
	}()
}
