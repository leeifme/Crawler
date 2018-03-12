package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
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

	utf8Reader := transform.NewReader(reps.Body,
		simplifiedchinese.GBK.NewDecoder())

	charset.DetermineEncoding()

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}
