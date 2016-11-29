package DatabaseConnector

import "testing"

func TestNewWorker(t *testing.T) {
	WorkerQueue := make(chan chan WorkRequest, 1)
	worker := NewWorker(1, WorkerQueue)
	if 1 != worker.ID {
		t.Errorf("Expected %d, was %d", 1, worker.ID)
	}

}

func TestStart(t *testing.T) {
	WorkerQueue := make(chan chan WorkRequest, 1)
	worker := NewWorker(1, WorkerQueue)
	if 1 != worker.ID {
		t.Errorf("Expected %d, was %d", 1, worker.ID)
	}
	WorkQueue := make(chan WorkRequest, 100)
	worker.Start()
	result := make(chan WorkResult)
	defer close(result)
	args := make([]interface{}, 1)
	WorkQueue <- WorkRequest{Query: "", Arguments: args, ResultChannel: result, F: func(w *WorkRequest) {
		w.ResultChannel <- WorkResult{Result: true, err: nil}
	}}
	var fail = <-result
	if fail.Result.(bool) != true {
		t.Errorf("Expected %t, was %+v", true, fail)
	}
	worker.Stop()

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
