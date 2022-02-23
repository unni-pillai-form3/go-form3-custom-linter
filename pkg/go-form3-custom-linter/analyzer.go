package go_form3_custom_linter

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "customlint",
	Doc:      "Checks for usages of custom code issues",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	expressionNodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.CallExpr)(nil),
	}
	pass.ResultOf[inspect.Analyzer].(*inspector.Inspector).Preorder(expressionNodeFilter, func(node ast.Node) {
		call := node.(*ast.CallExpr) // node is always a CallExpr thanks to nodeFilter

		// a sample check to identify usages of time.Sleep()
		if expr, ok := call.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := expr.X.(*ast.Ident); ok {
				if ident.Name == "time" && expr.Sel.Name == "Sleep" {
					pass.Reportf(node.Pos(), "no-sleep-lint")
				}
			}
		}

		// a sample check for identifying log.WithField usages with 'field' (1st) argument as non-constant
		isArgumentConstant(pass, node, call, "log", "WithField", 0, "no-log-withField-lint")

		return
	})

	return nil, nil
}

func isArgumentConstant(pass *analysis.Pass, node ast.Node, call *ast.CallExpr, identifier string, selector string, argumentPosition int, linterType string) {
	if expr, ok := call.Fun.(*ast.SelectorExpr); ok {
		if ident, ok := expr.X.(*ast.Ident); ok {
			if ident.Name == identifier && expr.Sel.Name == selector {
				args := call.Args
				if args != nil && len(args) > 0 {
					for i, argExpr := range args {
						if i == argumentPosition {
							if ident, ok := argExpr.(*ast.Ident); ok {
								if ident.Obj.Kind.String() != "const" {
									pass.Reportf(node.Pos(), linterType)
								}
							} else {
								pass.Reportf(node.Pos(), linterType)
							}
						}
					}
				}
			}
		}
	}
}
