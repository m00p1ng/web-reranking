package webgraph

import (
	"os"
	"path/filepath"
)

func isDirectory(p string) (bool, error) {
	fileInfo, err := os.Stat(p)
	return fileInfo.IsDir(), err
}

func pathTraverse(p string, h func(string)) {
	paths, err := filepath.Glob(p + "/*")
	if err != nil {
		panic(err)
	}
	for _, curpath := range paths {
		isDir, err := isDirectory(curpath)

		if err != nil {
			panic(err)
		}

		if isDir {
			pathTraverse(curpath, h)
		} else {
			h(curpath)
		}
	}
}
