package parser

import (
	"crawler/types"
	"regexp"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) types.ParseResult {
	ret := types.ParseResult{}
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)
	for _, match := range matchs {
		// ret.Items = append(ret.Items, string(match[2]))
		ret.Requests = append(ret.Requests, types.Request{
			URL:    string(match[1]),
			Parser: types.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return ret
}
