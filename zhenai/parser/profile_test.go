package parser

import (
	"io/ioutil"
	"testing"

	"../../engine"
	"../../model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "http://localhost:8080/mock/album.zhenai.com/u/6466817866957127625", "孤者何惧考莎啦")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element, but had %d\n", len(result.Items))
	}

	profile := result.Items[0]

	expected := engine.Item{
		Url:  "http://localhost:8080/mock/album.zhenai.com/u/6466817866957127625",
		Id:   "6466817866957127625",
		Type: "zhenai",
		PayLoad: model.Profile{
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

	if profile != expected {
		t.Errorf("Expected %v, but get %v\n", expected, profile)
	}
}
