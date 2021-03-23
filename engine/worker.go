package engine

import (
	"log"

	"crawler/fetcher"
	"crawler/types"
)

func Worker(r types.Request) (types.ParseResult, error) {
	log.Printf("Fetching %s\n", r.URL)
	content, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher error:\n\tfetching url: %s:%v\n", r.URL, err)
		return types.ParseResult{}, err
	}
	// parseResult := r.ParseFunc(content, r.URL)
	parseResult := r.Parser.Parser(content, r.URL)
	return parseResult, nil
}
