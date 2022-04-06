package main

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/chromedp/chromedp"
)

func FandomPage(id int) string {
	articleList, err := GetFandomArticleList()
	if err != nil {
		return fmt.Sprintf("Error %+v\n", err)
	}
	if !(0 <= id && id < len(articleList)) {
		return FandomIntError
	}

	article := articleList[id]
	url := fmt.Sprintf("%s%s", FandomHostname, article.Href)

	/*
		Read page via headless chrome
		document.querySelector("#mw-content-text > div").innerText
	*/

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var paragraphs string

	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate(`document.querySelector("#mw-content-text > div").innerText`, &paragraphs),
	)
	if err != nil {
		return fmt.Sprintf("headless chrome returned error %+v\n", err)
	}

	paragraphsOut := new(bytes.Buffer)

	for _, par := range strings.Split(paragraphs, "\n") {
		if strings.TrimSpace(par) != "" {
			fmt.Fprintln(paragraphsOut, GetWrappedText(strings.TrimSpace(par)))
		}
	}

	return paragraphsOut.String()

}
