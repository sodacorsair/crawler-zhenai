package main

import (
	"flag"
	"fmt"
	"github.com/Eru/distributedCrawler/rpcsupport"
	"github.com/Eru/distributedCrawler/worker"
	"log"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
