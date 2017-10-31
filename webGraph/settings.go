package main

import (
	"os"
	"path/filepath"

	"./webgraph"
)

var rootPath = filepath.Join(os.Getenv("HOME"), "Documents", "ir_proj", "html")

var tags = []webgraph.Tag{
	webgraph.Tag{Name: "a", Attribute: "href"},
}
