package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
)

var (
	ErrorUnexpectedErrorToken error             = errors.New("html.Error token unexpectedly recieved before the tokenizer found a <pre> block.")
	StoplistFromHrefCache     map[string]string = map[string]string{}
)

func GetStoplist(href string) (stoplist string, err error) {
	/*
		Given the link to the stoplist, grab
		the text (within the <pre> block)
	*/
	a, ok := StoplistFromHrefCache[href]
	if ok {
		return a, nil
	}

	resp, err := http.Get(href)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	tkn := html.NewTokenizer(bytes.NewReader(body))

	// token loop

	for {
		tt := tkn.Next()

		switch tt {
		case html.ErrorToken:
			return GetStoplistJS(href)

		case html.StartTagToken:
			/*
				Check if it is a pre block
			*/
			token := tkn.Token()
			if token.Data != "pre" {
				continue
			}

			/*
				Get the next token, which will be the text.
			*/
			tkn.Next()
			tokenNext := tkn.Token()
			data := tokenNext.Data
			StoplistFromHrefCache[href] = data
			return data, nil
		}
	}
}

func GetStoplistJS(href string) (string, error) {
	/*
		Derived from example at https://github.com/chromedp/examples/blob/master/eval/main.go
	*/

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(href),
		// chromedp.Evaluate(`document.querySelector("div#app.mr-1")`, &res),
		chromedp.Evaluate(`Application.data()`, &res),
	)
	if err != nil {
		return "", err
	}
	fmt.Println("res:", res, string(res))
	return string(res), err
}
