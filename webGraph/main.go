package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"./webgraph"
)

func main() {
	urlMap := webgraph.Urlmap(htmlPath)
	writeFile("urlmap.txt", func(file io.Writer) {
		for _, url := range urlMap {
			fmt.Fprintf(file, "%s\n", url)
		}
	})

	webgraph.GetGraph(urlMap, tags)
}

func writeFile(fn string, h func(io.Writer)) {
	cur, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	outputPath := filepath.Join(cur, "output")
	outputFile := filepath.Join(outputPath, fn)

	if isNotExist(outputPath) {
		os.Mkdir(outputPath, 0700)
	}

	file, err := os.Create(outputFile)

	if err != nil {
		panic(err)
	}

	h(file)
	file.Close()
}

func isNotExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}
