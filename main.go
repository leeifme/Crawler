package main

import (
	"leeif.me/Crawler/engine"
	"leeif.me/Crawler/zhenai/parser"
)

func main() {

	engine.SimpleEngine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
