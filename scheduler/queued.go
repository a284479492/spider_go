package scheduler

import "crawler/types"

type QueuedScheduler struct {
	requesterChan chan types.Request
	workerChan    chan chan types.Request
}

func (s *QueuedScheduler) Submit(r types.Request) {
	s.requesterChan <- r
}

func (s *QueuedScheduler) WorkerReady(r chan types.Request) {
	s.workerChan <- r
}

func (s *QueuedScheduler) Run() {
	s.requesterChan = make(chan types.Request)
	s.workerChan = make(chan chan types.Request)
	go func() {
		var rQ []types.Request
		var wQ []chan types.Request
		for {
			var activeRequest types.Request
			var activeWorker chan types.Request
			if len(rQ) >= 1 && len(wQ) >= 1 {
				activeRequest = rQ[0]
				activeWorker = wQ[0]
			}
			select {
			case requestQ := <-s.requesterChan:
				rQ = append(rQ, requestQ)
			case workerQ := <-s.workerChan:
				wQ = append(wQ, workerQ)
			case activeWorker <- activeRequest:
				rQ = rQ[1:]
				wQ = wQ[1:]
			}
		}
	}()
}
