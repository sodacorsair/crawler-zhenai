package main

import (
	"github.com/Eru/crawler/web/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("D:/Development/Go_Repo/go/src/github.com/Eru/crawler/web/resources/views")))
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"D:/Development/Go_Repo/go/src/github.com/Eru/crawler/web/resources/views/list.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
