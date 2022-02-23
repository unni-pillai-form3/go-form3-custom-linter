package main

import (
	go_form3_custom_linter "github.com/go-form3-custom-linter/pkg/go-form3-custom-linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(go_form3_custom_linter.Analyzer)
}
