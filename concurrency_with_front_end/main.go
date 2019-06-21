package main

import (
	"github.com/Eru/crawler/engine"
	"github.com/Eru/crawler/persist"
	"github.com/Eru/crawler/scheduler"
	"github.com/Eru/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		"dating_profile")
	if err != nil {
		panic(err)
	}

    e := engine.ConcurrentEngine{
    	Scheduler: &scheduler.QueuedScheduler{},
    	WorkerCount: 100,
    	ItemChan: itemChan,
	}

    e.Run(engine.Request{
    	Url: "https://www.zhenai.com/zhenghun",
    	ParseFunc: parser.ParseCityList,
	})
}
