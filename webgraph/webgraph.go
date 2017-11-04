package webgraph

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// Tag --
type Tag struct {
	Name      string
	Attribute string
}

// WebGraph --
type WebGraph struct {
	URLmap URLList
	OutURL [][]int
}

func extractURLFromFile(p string, t []Tag) URLList {
	file, err := os.Open(p)
	var urlOut URLList

	if err != nil {
		panic(err)
	}
	defer file.Close()

	urlOut = parseHTML(file, t)
	urlRedirect := extractRedirectLink(file)
	urlOut = append(urlOut, urlRedirect...)

	return urlOut
}

func getOutLinkURL(rp string, t []Tag) {
	pathTraverse(rp, func(curpath string) {
		curURL := splitRootURL(curpath, rp)
		fmt.Println("CURPATH => ", curURL)
		urlOut := extractURLFromFile(curpath, t)
		for i, u := range urlOut {
			uParse, err := url.Parse(u)

			if err != nil {
				panic(err)
			}

			if uParse.IsAbs() {
				urlOut[i] = removeHTTPPrefix(u)
			} else {
				urlOut[i] = strings.TrimRight(u, "/")
			}
		}
		urlOut.print()
		fmt.Println()
	})
}

func parseHTML(ct io.Reader, t []Tag) URLList {
	d := html.NewTokenizer(ct)
	var urlOut URLList

	for {
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := d.Token()
		switch tokenType {
		case html.StartTagToken:
			url := extractTagWithAttribute(token, t)
			if url != "" && !urlOut.include(url) {
				urlOut = append(urlOut, url)
			}
		}
	}
	return urlOut
}

func extractRedirectLink(ct io.Reader) URLList {
	pattern := `.*?window\.location\s*=\s*\"([^"]+)\"`
	re := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(ct)

	var urlRedirect URLList
	for scanner.Scan() {
		g := re.FindStringSubmatch(scanner.Text())
		if len(g) > 0 {
			urlRedirect = append(urlRedirect, g[1])
		}
	}
	return urlRedirect
}

func extractTagWithAttribute(tk html.Token, t []Tag) string {
	for _, tag := range t {
		isExpectTag := tk.Data == tag.Name
		if isExpectTag {
			for _, attr := range tk.Attr {
				if attr.Key == tag.Attribute {
					return attr.Val
				}
			}
		}
	}
	return ""
}

// GetGraph -- get webgraph
func GetGraph(urlMap URLList, t []Tag) {
}
