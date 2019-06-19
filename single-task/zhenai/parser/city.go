package parser

import (
	"github.com/Eru/crawler/engine"
	"regexp"
	"strconv"
)

const cityRe = `href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(
	contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	num := 0

	result := engine.ParseResult{}
	for _, m := range matches {
		// attention: in the closure, the anonymous function uses free variables when executing
		name := string(m[2])
		result.Items = append(
			result.Items, "User " + strconv.Itoa(num) + string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParseFunc: func(
					c []byte) engine.ParseResult {
						return ParseProfile(
							c, name)
				},
			})
		num++
	}

	return result
}
