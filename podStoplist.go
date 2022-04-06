package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"

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
			return "", ErrorUnexpectedErrorToken

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
