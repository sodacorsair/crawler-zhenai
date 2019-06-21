package parser

import (
	"github.com/Eru/crawler/engine"
	"github.com/Eru/crawler/model"
	"io/ioutil"
	"testing"
)

func TestProfile(t *testing.T) {
	// memory will boom if reading directly from website using fetcher
	contents, err := ioutil.ReadFile(
			"profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(
		contents, "http://album.zhenai.com/u/1375346415", "我是你的公主吗")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 " +
			"elements; but was %v", len(result.Items))
	}

	actual := result.Items[0]

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

	if actual != expected {
		t.Errorf("expected %v; but was %v",
			expected, actual)
	}

	//if profile.Name != expected.Name {
	//	t.Errorf("expected: %s; but was %s,",
	//		profile.Name, expected.Name)
	//}
	//if profile.Gender != expected.Gender {
	//	t.Errorf("expected: %s; but was %s,",
	//		profile.Gender, expected.Gender)
	//}
	//if profile.Age != expected.Age {
	//	t.Errorf("expected: %d; but was %d,",
	//		profile.Age, expected.Age)
	//}
	//if profile.Height != expected.Height {
	//	t.Errorf("expected: %d; but was %d,",
	//		profile.Height, expected.Height)
	//}
	//if profile.Weight != expected.Weight {
	//	t.Errorf("expected: %d; but was %d,",
	//		profile.Weight, expected.Weight)
	//}
	//if profile.Income != expected.Income {
	//	t.Errorf("expected: %s; but was %s,",
	//		profile.Income, expected.Income)
	//}
	//if profile.Marriage != expected.Marriage {
	//	t.Errorf("expected: %s; but was %s,",
	//		profile.Marriage, expected.Marriage)
	//}
	//if profile.Education != expected.Education {
	//	t.Errorf("expected: %s; but was %s,",
	//		profile.Education, expected.Education)
	//}
	//if profile.Occupation != expected.Occupation {
	//	t.Errorf("expected: %s; but was %s,",
	//		profile.Occupation, expected.Occupation)
	//}
	//if profile.Hukou != expected.Hukou {
	//	t.Errorf("expected: %s; but was %s,",
	//		profile.Hukou, expected.Hukou)
	//}
	//if profile.Constellation != expected.Constellation {
	//	t.Errorf("expected: %s; but was %s,",
	//		profile.Constellation, expected.Constellation)
	//}
	//if profile.House != expected.House {
	//	t.Errorf("expected: %s; but was %s,",
	//		profile.House, expected.House)
	//}
	//if profile.Car != expected.Car {
	//	t.Errorf("expected: %s; but was %s,",
	//		profile.Car, expected.Car)
	//}




}
