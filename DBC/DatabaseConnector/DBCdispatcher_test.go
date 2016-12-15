package DatabaseConnector

import "testing"

func TestStopWorker(t *testing.T) {
	workers = nil
	result := stopWorker()
	if result == true {
		t.Errorf("How could it stop the worker? there is no worker queue yet\nExpected false was: %+v", result)

	}
	startDispatcher(1)
	if stopWorker() != true {
		t.Errorf("Could not stop the worker!")
	}
	if stopWorker() == true {
		t.Errorf("How could it stop any worker? there are no workers anymore")
	}
	stopDispatcher()
	if stopWorker() == true {
		t.Errorf("How could it stop any worker? The worker queue is closed")
	}
}

func TestStopDispatcher(t *testing.T) {
	startDispatcher(3)
	stopDispatcher()
}

func TestStartDispatcher(t *testing.T) {
	startDispatcher(5)
	for stopWorker() {
	}
	stopDispatcher()
}
