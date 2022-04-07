package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/chromedp/chromedp"
	"github.com/gosimple/unidecode"
	"golang.org/x/net/html"
)

const (
	OrganStopHostname = "http://www.organstops.org"
	OrganStopNoArgs   = "Syntax: .stop <stop name>"
	OrganStopNotFound = "Organ stop not found."
)

func OrganStopHandler(argv []string) string {
	if len(argv) == 0 {
		return OrganStopNoArgs
	}

	stopname := unidecode.Unidecode(strings.Join(argv, " "))
	url, ok := OrganStops[stopname]
	if !ok {
		return OrganStopNotFound
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var text string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate(`document.querySelector("body").innerText`, &text),
	)
	if err != nil {
		return fmt.Sprintf("headless chrome error %+v\n", err)
	}

	buf := new(bytes.Buffer)

	for _, par := range strings.Split(text, "\n") {
		if strings.TrimSpace(par) != "" {
			fmt.Fprintln(buf, strings.TrimSpace(par))
		}
	}

	return text
}

var OrganStops map[string]string = map[string]string{}

func GetOrganStops(stops map[string]string) {
	url := fmt.Sprintf("%s/FullIndex.html", OrganStopHostname)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error during organ stop get: %+v\n", err)
		return
	}

	defer resp.Body.Close()
	tkn := html.NewTokenizer(resp.Body)
	//token loop
	for {
		tt := tkn.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			token := tkn.Token()
			if token.Data != "a" {
				continue
			}
			for _, attr := range token.Attr {
				if attr.Key == "href" && unicode.IsLetter(rune(attr.Val[0])) {
					href := fmt.Sprintf("%s/%s", OrganStopHostname, attr.Val)
					tkn.Next()
					token2 := tkn.Token()
					name := strings.ToLower(unidecode.Unidecode(token2.Data))
					stops[name] = href
				}
			}
		}
	}
}

func init() {
	go GetOrganStops(OrganStops)
}
