package webgraph

import (
	"fmt"
	"log"
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

	log.Println("Mapping URL...")
	pathTraverse(rp, func(curpath string) {
		url := urlmap(curpath, rp)
		urls = append(urls, url)
	})
	log.Println("URL mapped")

	return urls
}

func (ul URLList) print() {
	for _, u := range ul {
		fmt.Println(u)
	}
}

func (ul URLList) include(url string) bool {
	for _, u := range ul {
		if u == url {
			return true
		}
	}
	return false
}

func removeHTTPPrefix(url string) string {
	re := regexp.MustCompile(`^https?://(?:www\.)?(.*?)/?$`)
	g := re.FindStringSubmatch(url)

	if len(g) > 0 {
		return g[1]
	}
	return url
}

func joinURL(curPath string, relPath string) string {
	fullPath := ""

	return fullPath
}
