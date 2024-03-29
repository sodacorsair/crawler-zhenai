package engine

type Request struct {
	Url string
	ParseFunc
}

type ParseFunc func(
	contents []byte, url string) ParseResult

type ParseResult struct {
	Requests []Request
	Items []Item
}

type Item struct {
	Url string
	Type string
	Id string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
