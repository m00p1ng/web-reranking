package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"./webgraph"
)

func main() {
	wg := webgraph.GetGraph(htmlPath, tags)

	writeFile("urlmap.txt", func(file io.Writer) {
		for _, url := range wg.URLmap {
			fmt.Fprintf(file, "%s\n", url)
		}
	})

	writeFile("webgraph.txt", func(file io.Writer) {
		for _, out := range wg.OutURL {
			if len(out) == 0 {
				fmt.Fprintf(file, "-\n")
			} else {
				result := ""
				for _, it := range out {
					result += "," + strconv.Itoa(it)
				}
				result = strings.Trim(result, ",")
				fmt.Fprintf(file, "%s\n", result)
			}
		}
	})
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
	defer file.Close()

	h(file)
}

func isNotExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}
