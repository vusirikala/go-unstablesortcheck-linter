package main

import (
	"github.com/vusirikala/go-unstablesortcheck-linter/linter/unstablesortcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
