package main

import (
	"github.com/go-form3-custom-linter/pkg/src/constantarg"
	"github.com/go-form3-custom-linter/pkg/src/timeusage"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(constantarg.ConstantArgumentAnalyzer,
		timeusage.TimeAnalyzer)
}
