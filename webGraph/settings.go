package main

import (
	"os"
	"path/filepath"

	"./webgraph"
)

var htmlPath = filepath.Join(os.Getenv("HOME"), "Documents", "ir_proj", "html")

var tags = []webgraph.Tag{
	webgraph.Tag{Name: "a", Attribute: "href"},
}
