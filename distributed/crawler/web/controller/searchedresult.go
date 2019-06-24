package controller

import (
	"github.com/Eru/crawler/engine"
	"github.com/Eru/crawler/web/model"
	"github.com/Eru/crawler/web/view"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchedResultView
	client *elastic.Client
}

func CreateSearchResultHandler(
	template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view: view.CreateSearchedResultView(
			template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(
	w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
	}
	err = h.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(
	q string, from int) (model.SearchedResult, error) {
	var result model.SearchedResult
	result.Query = q
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(
			rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(
		reflect.TypeOf(engine.Item{}))
	result.PrevFrom =
		result.Start - len(result.Items)
	result.NextFrom =
		result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*:)`)
	return re.ReplaceAllString(q, "Payload.$1")
}
