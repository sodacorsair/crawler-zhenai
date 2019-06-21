package persist

import (
	"encoding/json"
	"github.com/Eru/crawler/engine"
	"github.com/Eru/crawler/model"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {

	expected := engine.Item{
		Url: "http://album.zhenai.com/u/1375346415",
		Type: "zhenai",
		Id: "1375346415",

		Payload: model.Profile{
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

	// TODO: Try to start up elastic search
	// here using docker to go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// save item
	err = save(client, index, expected)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index(index).
		Type(expected.Type).Id(expected.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(
		[]byte(*resp.Source), &actual)

    if err != nil {
    	panic(err)
	}

	actualProfile, err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v",
			actual, expected)
	}

}
