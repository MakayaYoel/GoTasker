package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/MakayaYoel/GoTasker/structs"
	"github.com/MakayaYoel/GoTasker/tasks"
	"github.com/joho/godotenv"
)

var (
	workerAmount int
	runTime      int
)

func init() {
	// Load Environment Variable
	if err := godotenv.Load(); err != nil {
		log.Fatal("Please create an environment file (.env) to use this program.")
	}

	workerAmountEnv, ok := os.LookupEnv("WORKER_AMOUNT")
	if !ok {
		log.Fatal("Please specify a WORKER_AMOUNT in your environment file.")
	}

	intWorkerAmountEnv, err := strconv.Atoi(workerAmountEnv)
	if err != nil {
		log.Fatalf("There was an error trying to convert the value of WORKER_AMOUNT into an int: %s", err.Error())
	}

	runTimeEnv, ok := os.LookupEnv("RUN_TIME")
	if !ok {
		log.Fatal("Please specify a RUN_TIME in your environment file.")
	}

	intRunTimeEnv, err := strconv.Atoi(runTimeEnv)
	if err != nil {
		log.Fatalf("There was an error trying to convert the value of RUN_TIME into an int: %s", err.Error())
	}

	workerAmount = intWorkerAmountEnv
	runTime = intRunTimeEnv
}

func main() {
	taskQueue := structs.NewTaskQueue()
	taskQueue.Start()

	for i := 0; i < workerAmount; i++ {
		worker := structs.NewTaskWorker(i, taskQueue)
		worker.Start()
	}

	// Send tasks for runTime seconds
	var tasksProcessed int
	startTime := time.Now()
	for {
		if time.Since(startTime).Seconds() >= float64(runTime) {
			break
		}

		switch rand.Intn(3) {
		case 0:
			taskQueue.SubmitTask(tasks.NewPrinterTask("Hello, World"))
		case 1:
			taskQueue.SubmitTask(tasks.NewSortingTask([]int{2, 4, 6, 7, 1, 9, 10, 5, 8, 3}))
		case 2:
			taskQueue.SubmitTask(tasks.NewAPIRequestTask())
		}

		tasksProcessed++
	}

	taskQueue.Wait()
	fmt.Printf("\nProcessed %d Tasks\n", tasksProcessed)
}
