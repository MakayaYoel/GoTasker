package tasks

import (
	"fmt"
	"time"
)

type apiRequestTask struct{}

// NewAPIRequestTask creates a new apiRequestTask struct.
func NewAPIRequestTask() *apiRequestTask {
	return &apiRequestTask{}
}

// Execute calls a cat fact API and prints the result.
func (aRT *apiRequestTask) Execute() error {
	time.Sleep(500 * time.Millisecond) // simulate API request delay
	fmt.Println("Finished API Request.")
	return nil
}
