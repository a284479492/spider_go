package view

import (
	"os"
	"testing"

	"crawler/frontend/model"
	common "crawler/model"
	"crawler/types"
)

func TestSearchResult_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template_test.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	page := model.SearchResult{}
	page.Hits = 123
	item := types.Item{
		Url:  "http://localhost:8080/mock/album.zhenai.com/u/6466817866957127625",
		Id:   "6466817866957127625",
		Type: "zhenai",
		PayLoad: common.Profile{
			Name:       "孤者何惧考莎啦",
			Gender:     "男",
			Age:        45,
			Weight:     266,
			Height:     181,
			Income:     "5001-8000元",
			Marriage:   "未婚",
			Education:  "博士及以上",
			Occupation: "人事/行政",
			Hokou:      "宁波市",
			Xinzuo:     "处女座",
			House:      "租房",
			Car:        "无车",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		t.Errorf("Render ERROR :%v\n", err)
	}
}
