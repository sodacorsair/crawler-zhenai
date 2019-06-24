package main

import (
	"github.com/Eru/crawler/engine"
	"github.com/Eru/crawler/persist"
	"github.com/Eru/crawler/scheduler"
	"github.com/Eru/crawler/zhenai/parser"
	"github.com/Eru/distributedCrawler/config"
)

func main() {
	itemChan, err := persist.ItemSaver(
		"dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "https://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList),
	})
}
