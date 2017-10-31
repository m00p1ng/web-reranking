package webgraph

import (
	"fmt"
	"regexp"
	"strings"
)

// URL -- url
type URL string

// URLList -- list of URL
type URLList []URL

func splitRootURL(p string, rp string) string {
	return strings.Replace(p, regexp.QuoteMeta(rp+"/"), "", -1)
}

func urlmap(p string, rp string) URL {
	var url = URL(splitRootURL(p, rp))
	return url
}

// Urlmap -- map directory to urlpath
func Urlmap(rp string) URLList {
	var urls URLList
	pathTraverse(rp, func(curpath string) {
		url := urlmap(curpath, rp)
		urls = append(urls, url)
	})
	return urls
}

func (u URLList) print() {
	for _, url := range u {
		fmt.Println(url)
	}
}
