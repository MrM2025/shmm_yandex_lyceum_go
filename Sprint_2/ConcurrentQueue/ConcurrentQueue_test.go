package main

import (
	"testing"
)

func TestConcurrentQueue_EnqueueDequeue(t *testing.T) {
	concurrentQueue := &ConcurrentQueue{}

	testCases := []struct {
		name     string
		elements []interface{}
	}{
		{
			name:     "Test Case 1",
			elements: []interface{}{1, 2, 3},
		},
		{
			name:     "Test Case 2",
			elements: []interface{}{"a", "b", "c"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			for _, element := range tc.elements {
				concurrentQueue.Enqueue(element)
			}

			for _, expected := range tc.elements {
				actual := concurrentQueue.Dequeue()
				if actual != expected {
					t.Errorf("Unexpected element. Got: %v, Expected: %v, %v", actual, expected, concurrentQueue.queue)
				}
	}})}}