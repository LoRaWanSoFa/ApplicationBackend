package DatabaseConnector

import "testing"

func TestNewWorker(t *testing.T) {
	WorkerQueue_test := make(chan chan WorkRequest, 1)
	worker := NewWorker(1, WorkerQueue_test)
	if 1 != worker.ID {
		t.Errorf("Expected %d, was %d", 1, worker.ID)
	}

}

func TestStart(t *testing.T) {
	// StartDispatcher(1)
	// worker := NewWorker(1, WorkerQueue)
	// if 1 != worker.ID {
	// 	t.Errorf("Expected %d, was %d", 1, worker.ID)
	// }
	// result := make(chan WorkResult)
	// defer close(result)
	// args := make([]interface{}, 1)
	// go func() {
	// 	WorkQueue <- WorkRequest{Query: "", Arguments: args, ResultChannel: result, F: func(w *WorkRequest) {
	// 		w.ResultChannel <- WorkResult{Result: true, err: nil}
	// 		result <- WorkResult{Result: true, err: nil}
	// 	}}
	// }()
	// fail := <-result
	// if fail.Result != true {
	// 	t.Errorf("Expected %t, was %+v", true, fail)
	// }
	// worker.Stop()

}

func TestStop(t *testing.T) {
	WorkerQueue := make(chan chan WorkRequest, 1)
	worker := NewWorker(1, WorkerQueue)
	if 1 != worker.ID {
		t.Errorf("Expected %d, was %d", 1, worker.ID)
	}
	worker.Stop()
	if 1 != worker.ID {
		t.Errorf("Expected %d, was %d", 1, worker.ID)
	}
}
