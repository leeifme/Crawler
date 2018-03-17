package parser

import (
	"regexp"

	"leeif.me/Crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: nil,
		})

		// fmt.Printf("City: %s , Url: %s\n", m[2], m[1])
	}
	// fmt.Printf("Matches find: %d\n", len(matches))
	return result
}
