package engine

import (
	"github.com/Eru/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(Url)
		if err != nil {
			log.Printf("Fetcher: error " +
				"fetching url %s: %v",
				Url, err)
			continue
		}

		parseResult := ParseFunc(body)
		requests = append(requests,
			Requests...)

		for _, item := range Items {
            log.Printf("Got item %v", item)
		}
	}
}
