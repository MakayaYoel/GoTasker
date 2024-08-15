package tasks

import "fmt"

type printerTask struct {
	message string
}

// NewPrinterTask creates a new printerTask struct with the given message.
func NewPrinterTask(message string) *printerTask {
	return &printerTask{
		message: message,
	}
}

// Execute sends a message.
func (pT *printerTask) Execute() error {
	fmt.Println(pT.message)
	return nil
}
