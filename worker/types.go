package worker

import (
	"errors"
	"fmt"
	"log"

	"crawler/config"
	"crawler/types"
	"crawler/zhenai/parser"
)

type SerializedParser struct {
	FunctionName string
	Args         interface{}
}

type Request struct {
	URL    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []types.Item
	Requests []Request
}

func SerializeRequest(r types.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		URL: r.URL,
		Parser: SerializedParser{
			FunctionName: name,
			Args:         args,
		},
	}
}

func SerializeResult(r types.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, request := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(request))
	}
	return result
}

func DeserializeRequest(r Request) (types.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return types.Request{}, err
	}
	return types.Request{
		URL:    r.URL,
		Parser: parser,
	}, nil
}

func deserializeParser(s SerializedParser) (types.Parser, error) {
	switch s.FunctionName {
	case config.ParseCity:
		return types.NewFuncParser(parser.ParseCity, s.FunctionName), nil
	case config.ParseCityList:
		return types.NewFuncParser(parser.ParseCityList, s.FunctionName), nil
	case config.ParseProfile:
		userName, ok := s.Args.(string)
		if !ok {
			return nil, fmt.Errorf("invalid args:%v", s.Args)
		} else {
			return parser.NewProfileParser(userName), nil
		}
	case config.NilParse:
		return types.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name: " + s.FunctionName)
	}
}

func DeserializeParseResult(r ParseResult) types.ParseResult {
	requests := []types.Request{}
	for _, request := range r.Requests {
		deserializeRequest, err := DeserializeRequest(request)
		if err != nil {
			log.Printf("error deserializing request %v: %v", request, err)
			continue
		}
		requests = append(requests, deserializeRequest)
	}
	return types.ParseResult{
		Items:    r.Items,
		Requests: requests,
	}
}
