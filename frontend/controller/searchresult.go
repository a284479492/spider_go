package controller

import (
	"context"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/olivere/elastic/v7"

	"crawler/frontend/model"
	"crawler/frontend/view"
	"crawler/types"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	var page model.SearchResult
	page, err = s.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = s.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (s SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := s.client.Search("dating_profile").Query(elastic.NewQueryStringQuery(rewriteQ(q))).From(from).Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Query = q
	// result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	for _, v := range resp.Each(reflect.TypeOf(types.Item{})) {
		result.Items = append(result.Items, v.(types.Item))
	}
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func rewriteQ(str string) string {
	re := regexp.MustCompile("([A-Z][a-z]*):")
	ans := re.ReplaceAll([]byte(str), []byte("PayLoad.$1:"))
	return string(ans)
}
