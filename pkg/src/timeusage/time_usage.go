package timeusage

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var TimeAnalyzer = &analysis.Analyzer{
	Name:     "timeusage",
	Doc:      "Checks for usages of time functions",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	//go analysis will populate the `pass.ResultOf` map with the prerequisite analyzers
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.CallExpr)(nil),
	}
	inspect.Preorder(nodeFilter, func(node ast.Node) {
		call := node.(*ast.CallExpr) // node is always a CallExpr thanks to nodeFilter

		// a sample check to identify usages of time.Sleep()
		doesFunctionCallExist(pass, node, call, "time", "Sleep")

		return
	})

	return nil, nil
}

func doesFunctionCallExist(pass *analysis.Pass, node ast.Node, call *ast.CallExpr, identifier string, selector string) {
	if expr, ok := call.Fun.(*ast.SelectorExpr); ok {
		if ident, ok := expr.X.(*ast.Ident); ok {
			if ident.Name == identifier && expr.Sel.Name == selector {
				pass.Reportf(node.Pos(), fmt.Sprintf("%s-%s", identifier, selector))
			}
		}
	}
}
