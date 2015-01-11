package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	url = "http://golang.org/ref/spec"
)

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBufferString("")
	doc.Find("pre.ebnf").Each(func(_ int, s *goquery.Selection) {
		fmt.Fprintln(buf, s.Text())
	})
}
