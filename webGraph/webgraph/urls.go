package webgraph

import (
	"fmt"
	"regexp"
	"strings"
)

// URLList -- list of URL
type URLList []string

func splitRootURL(path string, rootPath string) string {
	return strings.Replace(path, regexp.QuoteMeta(rootPath+"/"), "", -1)
}

func urlmap(path string, rootPath string) string {
	var url = splitRootURL(path, rootPath)
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
func (ul URLList) print() {
	for _, u := range ul {
		fmt.Println(u)
	}
}
