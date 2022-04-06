package main

/*
Organ title (builder, year)
document.querySelector("h3.organ-title").textContent

Location
document.querySelector("p.card-text").textContent

Technical details:
document.querySelector("#accordion-divisions").textContent
document.querySelector("#accordion-consoles").textContent
document.querySelector("#accordion-notes").textContent
*/

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/chromedp/chromedp"
)

func PodOrganHandler(id int) string {
	url := PodGetOrganUrl(id)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var title string
	var location string
	var divisions string
	var consoles string
	var notes string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate(`document.querySelector("h3.organ-title").innerText`, &title),
		chromedp.Evaluate(`document.querySelector("p.card-text").innerText`, &location),
		// chromedp.Evaluate(`document.querySelector("#accordion-divisions").textContent`, &divisions),
		chromedp.Evaluate(`document.querySelector("#accordion-divisions> div > div.collapse.item-1").className = "collapse.item-1.show"`, nil),
		chromedp.Evaluate(`document.querySelector("#accordion-divisions> div > div.collapse\\.item-1\\.show").innerText`, &divisions),
		// chromedp.Evaluate(`document.querySelector("#accordion-consoles").textContent`, &consoles),
		chromedp.Evaluate(`document.querySelector("#accordion-consoles > div > div.collapse.item-1").className = "collapse.item-1.show"`, nil),
		chromedp.Evaluate(`document.querySelector("#accordion-consoles > div > div.collapse\\.item-1\\.show").innerText`, &consoles),
		// chromedp.Evaluate(`document.querySelector("#accordion-notes").textContent`, &notes),
		chromedp.Evaluate(`document.querySelector("#accordion-notes> div > div.collapse.item-1").className = "collapse.item-1.show"`, nil),
		chromedp.Evaluate(`document.querySelector("#accordion-notes> div > div.collapse\\.item-1\\.show").innerText`, &notes),
	)

	if err != nil {
		return fmt.Sprintf("headless chrome returned error %+v\n", err)
	}

	buf := new(bytes.Buffer)
	var allBits []string = []string{
		title,
		location,
		divisions,
		consoles,
		notes,
	}

	for _, bit := range allBits {
		item := IgnoreSpaces(bit)
		fmt.Fprintln(buf, item)
		fmt.Fprintln(buf, ".")
	}

	fmt.Fprintln(buf, url)
	return buf.String()

}

func IgnoreSpaces(str string) string {
	s := strings.Split(string(str), "\n")
	var output []string
	for _, line := range s {
		if strings.TrimSpace(line) != "" {
			output = append(output, strings.TrimSpace(line))
		}
	}
	return strings.Join(output, "\n")
}
