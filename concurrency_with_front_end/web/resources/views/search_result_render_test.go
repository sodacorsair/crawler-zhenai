package views

import (
	"github.com/Eru/crawler/engine"
	common "github.com/Eru/crawler/model"
	"github.com/Eru/crawler/web/model"
	view2 "github.com/Eru/crawler/web/view"
	"os"
	"testing"
)

func TestSearchResultedViewRender(t *testing.T) {
	view := view2.CreateSearchedResultView(
		"list.html")

	out, err := os.Create("template_test.html")

	page := model.SearchedResult{}
	page.Hits = 123
	item := engine.Item{
		Url: "http://album.zhenai.com/u/1375346415",
		Type: "zhenai",
		Id: "1375346415",

		Payload: common.Profile{
			Name: "我是你的公主吗",
			Gender: "女",
			Age: 20,
			Height: 170,
			Weight: 55,
			Income: "3千以下",
			Marriage: "未婚",
			Education: "大学本科",
			Occupation: "在校学生",
			Hukou: "河南开封",
			Constellation: "天蝎座",
			Car: "已买车",
			House: "已购房",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
