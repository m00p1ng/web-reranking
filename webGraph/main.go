package main

import (
	"./webgraph"
)

func main() {
	urlMap := webgraph.Urlmap(htmlPath)
	webgraph.GetGraph(urlMap, tags)
}
