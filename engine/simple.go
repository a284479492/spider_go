package engine

import (
	"crawler/types"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...types.Request) {
	requests := []types.Request{}
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item: %v\n", item)
		}
	}
}
