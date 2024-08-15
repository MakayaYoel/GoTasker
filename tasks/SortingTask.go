package tasks

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type sortingTask[T constraints.Ordered] struct {
	dataset []T
}

// NewSortingTask creates a new sortingTask struct with the given dataset.
func NewSortingTask[T constraints.Ordered](dataset []T) *sortingTask[T] {
	return &sortingTask[T]{
		dataset: dataset,
	}
}

// Execute sorts the dataset using bubble sort.
func (sT *sortingTask[T]) Execute() error {
	for i := 0; i < len(sT.dataset)-1; i++ {
		for j := 0; j < len(sT.dataset)-1-i; j++ {
			if sT.dataset[j] > sT.dataset[j+1] {
				sT.dataset[j], sT.dataset[j+1] = sT.dataset[j+1], sT.dataset[j]
			}
		}
	}

	fmt.Printf("Successfully sorted dataset: %v", sT.dataset)

	return nil
}
