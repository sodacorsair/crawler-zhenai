package main

import (
	"flag"
	"github.com/Eru/crawler/engine"
	"github.com/Eru/crawler/scheduler"
	"github.com/Eru/crawler/zhenai/parser"
	"github.com/Eru/distributedCrawler/config"
	itemsaver "github.com/Eru/distributedCrawler/persist/client"
	"github.com/Eru/distributedCrawler/rpcsupport"
	worker "github.com/Eru/distributedCrawler/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_hosts", "",
		"worker hosts (comma seperated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(
		*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(
		strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url: "https://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}

func createClientPool(
	hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewCilent(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v",
				h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
