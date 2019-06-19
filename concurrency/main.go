package main

import (
	"crawler-zhenai/single-task/zhenai/parser"
	"github.com/Eru/crawler/engine"
	"github.com/Eru/crawler/scheduler"
)

func main() {
    e := engine.ConcurrentEngine{
    	Scheduler: &scheduler.QueuedScheduler{},
    	WorkerCount: 100,
	}

    //e.Run(engine.Request{
    //	Url: "https://www.zhenai.com/zhenghun",
    //	ParseFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/shanghai",
		ParseFunc: parser.ParseCity,
	})
}
