# GoTasker

GoTasker is a task queue/worker pool implementation in Golang, designed to handle and process multiple tasks such as printing messages, sorting datasets, and making API requests utilizing Golang's concurrency.

## Features
- **Concurrent Task Processing:** Utilizes Go's goroutines to execute tasks concurrently across many workers.
- **Task Variety:** Processes different types of tasks, including `PrinterTask`, `SortingTask`, and `APIRequestTask`.
- **Scalability:** Easily scales by adjusting the number of workers in the environment file.

## Installation & Usage
To use this project, first make sure you have [Golang](https://go.dev/) installed on your system. Then, clone the repository onto your system and navigate to the project:
```bash
git clone https://github.com/MakayaYoel/GoTasker.git
cd GoTasker
```

Before running the project, make sure to rename the `.env.example` file to `.env`, and adjust the worker amount in the file to your liking.

Finally, run the project using the following command:
```bash
go run main.go
```