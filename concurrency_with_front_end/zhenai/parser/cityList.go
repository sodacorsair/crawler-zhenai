package parser

import (
	"github.com/Eru/crawler/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(
	contents []byte, _ string) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
        result.Requests = append(
        	result.Requests, engine.Request{
        		Url: string(m[1]),
        		ParseFunc: ParseCity,
			})
	}

    return result
}
