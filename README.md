# crawler-zhenai

A crawler used to retrieve posters' information from spouse-seeking website zhenai

### about single task version

So far the website has some registering windows bouncing, the original http.Get method has to be modified so that web pages could be visited normally.

### to-do

- Use css query selector or xpath to crawl more webpages
- Chanllege the anti-webspider technologies and observe robots protocols
- Simulate login to crawl active webpages
- Distributed deduplication (use RPC to package)
- Optimize ElasticSearch querying quality
- Add dropdown searching window
- Crawl and display albums
- Face rank or word segmentation or any other NLP, big data and AI toolkits
- Represent data with ElasticStack
- Shell, powershell or bat scripts to deploy
- Docker and Kubernetes deployment
- Integrate with service discovery framework like consul/etcd/zookeeper
- Assemble and analyze logs with Logstash

### acknowledgement

https://coding.imooc.com/class/180.html
