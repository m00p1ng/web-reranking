package webgraph

import (
	"bufio"
	"io"
	"log"
	"net/url"
	"os"
	"path"
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

func extractURLFromFile(path string, t []Tag) URLList {
	file, err := os.Open(path)
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

func getOutLinkURL(htmlPath string, url string, tag []Tag) URLList {
	contentPath := path.Join(htmlPath, url)
	log.Println("Reading...", contentPath)

	urlOut := extractURLFromFile(contentPath, tag)
	for i, u := range urlOut {
		rel, err := pathResolve(u, url)
		if err == nil {
			urlOut[i] = rel
		}
	}
	return urlOut
}

func pathResolve(u string, curPath string) (string, error) {
	uParse, err := url.Parse(u)

	if err != nil {
		return "", err
	}

	if uParse.IsAbs() {
		return removeHTTPPrefix(u), nil
	}

	if u != "/" {
		u = strings.TrimRight(u, "/")
	}

	return joinURL(curPath, u), nil
}

func parseHTML(content io.Reader, t []Tag) URLList {
	d := html.NewTokenizer(content)
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
			if url != "" && urlOut.Find(url) == -1 {
				urlOut = append(urlOut, url)
			}
		}
	}
	return urlOut
}

func extractRedirectLink(content io.Reader) URLList {
	pattern := `.*?window\.location\s*=\s*\"([^"]+)\"`
	re := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(content)

	var urlRedirect URLList
	for scanner.Scan() {
		g := re.FindStringSubmatch(scanner.Text())
		if len(g) > 0 {
			urlRedirect = append(urlRedirect, g[1])
		}
	}
	return urlRedirect
}

func extractTagWithAttribute(token html.Token, t []Tag) string {
	for _, tag := range t {
		isExpectTag := token.Data == tag.Name
		if isExpectTag {
			for _, attr := range token.Attr {
				if attr.Key == tag.Attribute {
					return attr.Val
				}
			}
		}
	}
	return ""
}

// GetGraph -- get webgraph
func GetGraph(path string, t []Tag) WebGraph {
	urlMap := Urlmap(path)

	log.Println("Creating Webgraph...")
	wg := WebGraph{
		URLmap: urlMap,
		OutURL: make([][]int, len(urlMap)),
	}

	for i, um := range wg.URLmap {
		urlOut := getOutLinkURL(path, um, t)

		for _, uo := range urlOut {
			idx := wg.URLmap.Find(uo)

			if idx != -1 {
				wg.OutURL[i] = append(wg.OutURL[i], idx+1)
			}
		}
	}
	log.Println("Webgraph created")

	return wg
}
