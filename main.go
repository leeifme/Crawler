package main

import (
	"leeif.me/Crawler/engine"
	"leeif.me/Crawler/scheduler"
	"leeif.me/Crawler/zhenai/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
