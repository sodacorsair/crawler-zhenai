package parser

import (
	"github.com/Eru/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(
	contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
        result.Items = append(
        	result.Items, "City " + string(m[2]))
        result.Requests = append(
        	result.Requests, engine.Request{
        		Url: string(m[1]),
        		ParseFunc: ParseCity,
			})
	}

    return result
}
