package DatabaseConnector

var WorkerQueue chan chan WorkRequest
var workers chan Worker
var stop bool

func StopWorker() bool {
	select {
	case w, ok := <-workers:
		if ok {
			w.Stop()
			return true
		} else {
			//channel workers is nil / not initialized
			return false
		}
	default:
		//No workers left
		return false
	}
}

func StopDispatcher() {
	defer close(workers)
	for StopWorker() {
	}
	stop = true
}

func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to put the workers' work channels into.
	WorkerQueue = make(chan chan WorkRequest, nworkers)
	workers = make(chan Worker, nworkers)
	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		//fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
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
			}
		}
	}()
}
