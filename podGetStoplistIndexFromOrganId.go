package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var StoplistWrongNumberOfStoplistHref error = errors.New("Didn't recieve any href to stoplist, or recieved more than one.")

func podGetStoplistIndexFromOrganId(organId int) (string, error) {
	organUrl := PodGetOrganUrl(organId)
	log.Println("request", organUrl)

	resp, err := http.Get(organUrl)
	if err != nil {
		return "", err
	}
	log.Println("^ request done")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	tkn := html.NewTokenizer(bytes.NewReader(body))

	var stoplistHref []string

	log.Println("start tokenizer")

	for {
		tt := tkn.Next()

		switch tt {
		case html.ErrorToken:
			// done
			log.Println("tokenizer done")
			if len(stoplistHref) != 1 {
				/*
					Something went wrong
				*/
				return "", StoplistWrongNumberOfStoplistHref
			}

			return fmt.Sprintf("%s%s", PodHostname, stoplistHref[0]), nil

		case html.StartTagToken:
			token := tkn.Token()
			if token.Data != "a" {
				continue
			}

			var attrs []html.Attribute = token.Attr
			for _, attr := range attrs {
				if attr.Key == "href" && strings.Contains(attr.Val, "/stoplist/") {
					stoplistHref = append(stoplistHref, attr.Val)
				}
			}

		}
	}

}
