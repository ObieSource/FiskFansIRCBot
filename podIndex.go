package main

import (
	"bufio"
	_ "embed"
	"strings"

	"github.com/gosimple/unidecode"
)

//go:embed "podIndex.csv"
var podIndex string

func PodKeywordSearch(words []string) (results []string) {
	read := strings.NewReader(podIndex)
	readscan := bufio.NewScanner(read)

	for {
		if readscan.Scan() == false {
			// reached end of file or error
			return
		}
		line := readscan.Text()

		var doesntContain bool = false
		for _, keyword := range words {
			if !strings.Contains(
				unidecode.Unidecode(strings.ToLower(line)),
				unidecode.Unidecode(strings.ToLower(keyword))) {
				doesntContain = true
				break
			}
		}
		if !doesntContain {
			// all keywords check out.
			results = append(results, strings.TrimSpace(line))
		}

	}
}
