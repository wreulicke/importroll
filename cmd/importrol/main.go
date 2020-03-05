package main

import (
	"github.com/wreulicke/importroll"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(importroll.Analyzer)
}
