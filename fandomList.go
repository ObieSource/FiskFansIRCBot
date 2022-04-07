package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type FandomArticle struct {
	Name string
	Href string
}

const (
	FandomHostname = "https://pipe-organ.fandom.com"
)

func GetFandomArticleList() (articles []FandomArticle, err error) {
	fullListUrl := fmt.Sprintf("%s/wiki/Special:AllPages", FandomHostname)

	resp, err := http.Get(fullListUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	/*
		Run html parser to find all
		a links with href starts with /wiki/
	*/
	err = nil
	tkn := html.NewTokenizer(resp.Body)
	var href string
	var title string
	for {
		tt := tkn.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			token := tkn.Token()
			if token.Data == "a" {
				// check for href

				href = ""
				title = ""
				for _, attr := range token.Attr {
					if attr.Key == "href" && strings.HasPrefix(attr.Val, "/wiki/") && !strings.HasPrefix(attr.Val, "/wiki/Special") {
						href = attr.Val
					} else if attr.Key == "title" {
						title = attr.Val
					}
				}
				if href != "" && title != "" {
					articles = append(articles, FandomArticle{
						title,
						href,
					})
				}
			}
		}

	}
}

func GetFandomArticlePrint() string {
	buf := new(bytes.Buffer)
	resp, err := GetFandomArticleList()
	if err != nil {
		return fmt.Sprintf("Error %+v\n", err)
	}
	for i, c := range resp {
		fmt.Fprintf(buf, "%2d - %s %s\n", i, c.Name, c.Href)
	}
	return buf.String()
}
