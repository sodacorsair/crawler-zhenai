package engine

var visited = make(map[string]bool)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
    ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i:= 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(),
			out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		for _, item := range result.Items {
            go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func isDuplicate(url string) bool {
	if visited[url] {
		return true
	}

	visited[url] = true
	return false
}

func createWorker(in chan Request,
	out chan ParseResult, ready ReadyNotifier) {
		go func() {
			for {
                ready.WorkerReady(in)
				request := <- in
				result, err := worker(request)
				if err != nil {
					continue
				}
				out <- result
			}
		}()
}