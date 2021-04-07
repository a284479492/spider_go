package engine

import (
	dupfilter "crawler/dupFilter"
	"crawler/types"
)

type ConcurrentEngine struct {
	Scheduler      Scheduler
	WorkerCount    int
	ItemChan       chan types.Item
	ProcessRequest Processor
}

type Processor func(types.Request) (types.ParseResult, error)

type Scheduler interface {
	Submit(types.Request)
	WorkerReady(r chan types.Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...types.Request) {
	out := make(chan types.ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.CreateWorker(out, e.Scheduler)
	}

	filter := &dupfilter.Filter{}
	filter.Run()

	//将请求发送到过滤器，进行去重操作
	for _, seed := range seeds {
		filter.CheckIn(seed)
	}

	//从过滤器里取出已经去重过的结果，然后提交给调度器，重新进入到任务分配工作
	go func() {
		for {
			e.Scheduler.Submit(filter.GetOut())
		}
	}()

	for {
		result := <-out

		for _, item := range result.Items {
			item0 := item
			go func() {
				e.ItemChan <- item0
			}()
		}

		//将请求发送到过滤器，进行去重操作
		for _, request := range result.Requests {
			filter.CheckIn(request)
		}
	}
}

func (e ConcurrentEngine) CreateWorker(out chan types.ParseResult, s Scheduler) {
	in := make(chan types.Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := e.ProcessRequest(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
