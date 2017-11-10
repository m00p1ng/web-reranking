# Web Reranking

## Web Graph and PageRank
Implement by GO

### How to install
```
$ git clone https://github.com/m00p1ng/web-reranking web-reranking
$ cd web-reranking
$ go get ./..
```

### How to run
```
$ cd web-reranking/webGraph
$ go run main.go settings.go

OR 

$ make run
```

### How to build
```
$ go bulid -o webgraph.out main.go settings.go 
$ ./webgraph.out

OR 

$ make build
```

## Indexing and Searching
Implement by lucene (Java)

**Index Files**
```
args
    -index <INDEX_PATH>
    -docs <DOCS_PATH>
    -pagerank <PAGERANK_PATH>
    -urlmap <URLMAP_PATH>
```

**Search Files**
```
args
    -index <INDEX_PATH>
    -docs <DOCS_PATH>
```