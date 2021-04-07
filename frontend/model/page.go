package model

import "crawler/types"

type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []types.Item
}
