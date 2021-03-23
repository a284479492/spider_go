package dupfilter

import (
	"crawler/types"
)

type Filter struct {
	in  chan types.Request
	out chan types.Request
}

func (f *Filter) CheckIn(req types.Request) {
	f.in <- req
}

func (f *Filter) GetOut() types.Request {
	ans := <-f.out
	return ans
}

func (f *Filter) Run() {
	f.in = make(chan types.Request)
	f.out = make(chan types.Request)
	go func() {
		for {
			req := <-f.in
			if !isDuplicate(req.URL) {
				f.out <- req
			}
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if _, has := visitedUrls[url]; has {
		return true
	}
	visitedUrls[url] = true
	return false
}
