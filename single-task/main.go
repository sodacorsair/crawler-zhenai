package single_task

import (
	"github.com/Eru/crawler/engine"
	"github.com/Eru/crawler/zhenai/parser"
)

func main() {
    engine.Run(engine.Request{
    	Url: "https://www.zhenai.com/zhenghun",
    	ParseFunc: parser.ParseCityList,
	})

}
