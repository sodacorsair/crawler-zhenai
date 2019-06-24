package main

import (
	"fmt"
	"github.com/Eru/distributedCrawler/config"
	"github.com/Eru/distributedCrawler/rpcsupport"
	"github.com/Eru/distributedCrawler/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewCilent(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1375346415",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "我是你的公主吗",
		},
	}

	var result worker.ParseResult
	err = client.Call(
		config.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
