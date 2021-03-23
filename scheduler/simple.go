package scheduler

import "crawler/types"

type SimpleScheduler struct {
	workerChan chan types.Request
}

func (s *SimpleScheduler) Submit(r types.Request) {
	go func() { s.workerChan <- r }()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan types.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) WorkerReady(c chan types.Request) {

}
func (s *SimpleScheduler) Run() {
}
