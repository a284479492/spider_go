package parser

import (
	"regexp"

	"crawler/types"
)

var (
	cityRe     = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityNextRe = regexp.MustCompile(`<span class="pager"><a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/.+)">([\d]+)</a></span>`)
)

func ParseCity(contents []byte, _ string) types.ParseResult {
	ret := types.ParseResult{}
	matchs := cityRe.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		// ret.Items = append(ret.Items, string(match[2]))
		ret.Requests = append(ret.Requests, types.Request{
			URL:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}
	cityMatchs := cityNextRe.FindAllSubmatch(contents, -1)
	for i := range cityMatchs {
		match := cityMatchs[i]
		ret.Requests = append(ret.Requests, types.Request{
			URL:    string(match[1]),
			Parser: types.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return ret
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parser(c []byte, url string) types.ParseResult {
	return ParseProfile(c, url, p.userName)
}
func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ParseProfile", p.userName
}

func NewProfileParser(userName string) *ProfileParser {
	return &ProfileParser{userName: userName}
}
