package webgraph

import (
    "bufio"
    "io"
    "log"
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
type WebGraph [][]int

func extractURLFromFile(path string, t []Tag) URLList {
    var urlOut URLList

    file, err := os.Open(path)
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
func GetGraph(path string, t []Tag) (URLList, WebGraph) {
    urlMap := Urlmap(path)
    wg := make(WebGraph, len(urlMap))

    log.Println("Creating Webgraph...")
    for i, um := range urlMap {
        urlOut := getOutLinkURL(path, um, t)

        for _, uo := range urlOut {
            idx := urlMap.Find(uo)

            if idx != -1 {
                wg[i] = append(wg[i], idx+1)
            }
        }
    }
    log.Println("Webgraph created")

    return urlMap, wg
}

func ReadGraph(p string) WebGraph {
    var wg WebGraph

    file, err := os.Open(p)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        if line == "-" {
            wg = append(wg, make([]int, 0))
        } else {
            num := strings.Split(line, ",")
            wg = append(wg, parseStringToInt(num))
        }
    }
    return wg
}
