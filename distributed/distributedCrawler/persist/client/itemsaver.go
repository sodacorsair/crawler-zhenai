package client

import (
	"github.com/Eru/crawler/engine"
	"github.com/Eru/distributedCrawler/config"
	"github.com/Eru/distributedCrawler/rpcsupport"
	"log"
)

func ItemSaver(
	host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewCilent(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			itemCount++

			// Call RPC to save item
			result := ""
			err := client.Call(config.ItemSaverRpc,
				item, &result)
			if err != nil {
				log.Printf("Item Saver: error "+"saving item %v: %v",
					item, err)
			}
		}
	}()
	return out, nil
}
