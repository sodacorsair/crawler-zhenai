package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(500 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//<- rateLimiter
	// handle bouncing window
    client := &http.Client{}
    request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
    request.Header.Add("Cookie", "xxxxxx")
    request.Header.Add("User-Agent", "xxx")
    request.Header.Add("X-Requested-With", "xxxx")

    resp, err := client.Do(request)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
        return nil,
        	fmt.Errorf("wrong status code: %d",
        		resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,
		e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(
	r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
        return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}
