package parser

import (
	"github.com/Eru/crawler/engine"
	"github.com/Eru/distributedCrawler/config"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(
		`href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityRe = regexp.MustCompile(
		`href="(http://m.zhenai.com/zhenghun/[^"]+)">`)
)

func ParseCity(
	contents []byte, url string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		// attention: in the closure, the anonymous function uses free variables when executing
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    string(m[1]),
				Parser: NewProfileParser(string(m[2])),
			})
	}

	matches = cityRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCity, config.ParseCity),
			})
	}

	return result
}
