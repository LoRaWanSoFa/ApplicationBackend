package DatabaseConnector

import "testing"

func TestStopWorker(t *testing.T) {
	workers = nil
	result := StopWorker()
	if result == true {
		t.Errorf("How could it stop the worker? there is no worker queue yet\nExpected false was: %+v", result)

	}
	StartDispatcher(1)
	if StopWorker() != true {
		t.Errorf("Could not stop the worker!")
	}
	if StopWorker() == true {
		t.Errorf("How could it stop any worker? there are no workers anymore")
	}
	StopDispatcher()
	if StopWorker() == true {
		t.Errorf("How could it stop any worker? The worker queue is closed")
	}
}

func TestStopDispatcher(t *testing.T) {
	StartDispatcher(3)
	StopDispatcher()
}

func TestStartDispatcher(t *testing.T) {
	StartDispatcher(5)
	for StopWorker() {
	}
	StopDispatcher()
}
