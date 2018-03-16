package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {

	reps, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer reps.Body.Close()

	if reps.StatusCode != http.StatusOK {
		fmt.Println("Error: Status Code: ", reps.StatusCode)
		return
	}

	e := determineEncoding(reps.Body)
	utf8Reader := transform.NewReader(reps.Body,
		e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	printCityList(all)
	// fmt.Printf("%s\n", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {

	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s , Url: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches find: %d\n", len(matches))
}
