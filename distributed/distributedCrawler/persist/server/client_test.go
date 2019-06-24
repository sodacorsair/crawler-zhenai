package main

import (
	"github.com/Eru/crawler/engine"
	"github.com/Eru/crawler/model"
	"github.com/Eru/distributedCrawler/config"
	"github.com/Eru/distributedCrawler/rpcsupport"
	"testing"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	// start ItemSaverServce
	//go serveRpc(host, "test1")
	//time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewCilent(host)
	if err != nil {
		panic(err)
	}

	// call service
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1375346415",
		Type: "zhenai",
		Id:   "1375346415",

		Payload: model.Profile{
			Name:          "我是你的公主吗",
			Gender:        "女",
			Age:           20,
			Height:        170,
			Weight:        55,
			Income:        "3千以下",
			Marriage:      "未婚",
			Education:     "大学本科",
			Occupation:    "在校学生",
			Hukou:         "河南开封",
			Constellation: "天蝎座",
			Car:           "已买车",
			House:         "已购房",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc,
		item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s",
			result, err)
	}
}
