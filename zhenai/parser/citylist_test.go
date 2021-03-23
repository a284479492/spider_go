package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents, "")

	const contentSize = 470
	if len(result.Requests) != contentSize {
		t.Errorf("Results should have %d requests; but had %d \n", contentSize, len(result.Requests))
	}

	if len(result.Items) != contentSize {
		t.Errorf("Results should have %d items; but had %d \n", contentSize, len(result.Items))
	}
}
