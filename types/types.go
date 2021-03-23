package types

type Request struct {
	URL    string
	Parser Parser
}

type Parser interface {
	Parser(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type ParseFunc func(contents []byte, url string) ParseResult

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	PayLoad interface{}
}

func NilParseFunc([]byte) ParseResult {
	return ParseResult{}
}

type NilParser struct {
}

func (p NilParser) Parser(c []byte, url string) ParseResult {
	return ParseResult{}
}
func (p NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParseFunc
	name   string
}

func (f *FuncParser) Parser(c []byte, url string) ParseResult {
	return f.parser(c, url)
}
func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
