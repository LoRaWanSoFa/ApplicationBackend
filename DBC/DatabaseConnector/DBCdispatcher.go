package DatabaseConnector

import "time"

//WorkerQueue contains a channel with WorkRequests
var WorkerQueue chan chan WorkRequest
var workers chan Worker
var stop bool

func stopWorker() bool {
	select {
	case w, ok := <-workers:
		if ok {
			w.stop()
			return true
		}
		//channel workers is nil / not initialized
		return false
	default:
		//No workers left
		return false
	}
}

func stopDispatcher() {
	defer close(workers)
	stop = true
	for stopWorker() {
	}
}

func startDispatcher(nworkers int) {
	// First, initialize the channel we are going to put the workers' work channels into.
	WorkerQueue = make(chan chan WorkRequest, nworkers)
	workers = make(chan Worker, nworkers)
	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		//fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.start()
		workers <- worker
	}

	stop = false

	go func() {
		for {

			select {
			case work := <-WorkQueue:
				//fmt.Println("Received work requeust")
				go func() {
					worker := <-WorkerQueue
					//fmt.Println("Dispatching work request")
					worker <- work
				}()
			default:
				if stop == true {
					break
				}
				time.Sleep(10 * time.Millisecond) //Sleep a bit when there is no work. performance reasons.
			}
		}
	}()
}
