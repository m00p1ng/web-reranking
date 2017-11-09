package webgraph

import (
    "fmt"
    "log"
    "net/url"
    "path"
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

// Print -- Print URLs
func (ul URLList) Print() {
    for _, u := range ul {
        fmt.Println(u)
    }
}

// Find -- Find URL and return index that found, is not return -1
func (ul URLList) Find(url string) int {
    for i, u := range ul {
        if u == url {
            return i
        }
    }
    return -1
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
    if string(relPath[0]) != "/" {
        base := path.Dir(curPath)
        fullPath := path.Join(base, relPath)
        return fullPath
    }

    cur, err := url.Parse("http://" + curPath)

    if err != nil {
        panic(err)
    }
    fullPath := path.Join(cur.Host, relPath)

    return fullPath
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
