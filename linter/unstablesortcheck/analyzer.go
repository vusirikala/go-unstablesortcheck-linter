package analyzer

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name: "unstablesortcheck",
	Doc:  "reports uses of sort.Sort and sort.Slice",
	Run:  run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

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
			if pkg.Name == "sort" && fn.Sel.Name == "Sort" {
				pass.Reportf(call.Pos(), "Use of sort.Sort. Replace it with sort.Stable")
			}
			if pkg.Name == "sort" && fn.Sel.Name == "Slice" {
				pass.Reportf(call.Pos(), "Use of sort.Slice. Replace it with sort.SliceStable")
			}
			return true
		})
	}
	return nil, nil
}
