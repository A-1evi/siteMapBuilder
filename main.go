package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	link "github.com/A-1evi/htmlLinkParser"
)

/*
1. GET the webpage
2. Parse all the links of thea page
3. build the proper urls with our links
4. remove all the other links that are not on the same webpage like github etc
5. find all pages(BFS)
6. print the xml page
*/
func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "The websiteh that we gonnna  use for site map")
	flag.Parse()

	pages := get(*urlFlag)
	for _, page := range pages {
		fmt.Println(page)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filterLinks(hrefs(resp.Body, base), hasPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)

		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filterLinks(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, l := range links {
		if keepFn(l) {
			ret = append(ret, l)
		}

	}
	return ret
}

func hasPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
