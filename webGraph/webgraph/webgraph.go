package webgraph

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

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

	if err == nil {
		urlOut = parseHTML(file, t)
		urlRedirect := extractRedirectLink(file)
		urlOut = append(urlOut, urlRedirect...)
	} else {
		panic(err)
	}
	defer file.Close()

	return urlOut
}

func getOutLinkURL(rp string, t []Tag) {
	pathTraverse(rp, func(curpath string) {
		curURL := splitRootURL(curpath, rp)
		fmt.Println("CURPATH => ", curURL)
		urlOut := extractURLFromFile(curpath, t)
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
			if url != "" {
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
		g := re.FindAllStringSubmatch(scanner.Text(), 1)
		for i := 0; i < len(g); i++ {
			urlRedirect = append(urlRedirect, g[i][1])
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
func GetGraph(rp string, t []Tag) {
	log.Println("Mapping URL...")
	Urlmap(rp)
	log.Println("URL mapped")
}
