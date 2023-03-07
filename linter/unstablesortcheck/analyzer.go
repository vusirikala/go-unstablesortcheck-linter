package analyzer

import (
	"errors"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "unstablesortcheck",
	Doc:  "reports uses of sort.Sort and sort.Slice",
	Run:  run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var ErrNoError = errors.New("no error")

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			fn, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			pkg, ok := fn.X.(*ast.Ident)
			if !ok {
				return true
			}
			if pkg.Name == "sort" && (fn.Sel.Name == "Sort" || fn.Sel.Name == "Slice") {
				pass.Reportf(call.Pos(), "use of %s.%s", pkg.Name, fn.Sel.Name)
			}
			return true
		})
	}
	return nil, ErrNoError
}
