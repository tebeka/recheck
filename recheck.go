package recheck

import (
	"fmt"
	"go/ast"
	"regexp"
	"regexp/syntax"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "recheck",
	Doc:  "Validate regular expressions",
	Run:  run,
}

var regexpFuncs = map[string]bool{
	"regexp.MustCompile":      true,
	"regexp.Compile":          true,
	"regexp.MustCompilePOSIX": true,
	"regexp.CompilePOSIX":     true,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			ce, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			funcName := nodeFuncName(ce)
			if !regexpFuncs[funcName] {
				return true
			}

			val, ok := ce.Args[0].(*ast.BasicLit)
			if !ok {
				return true
			}

			// Trim enclosing "" or ``
			expr := strconv.Quote(val.Value[1 : len(val.Value)-2])
			var check func(string) (*regexp.Regexp, error)
			if strings.HasSuffix(funcName, "POSIX") {
				check = regexp.CompilePOSIX
			} else {
				check = regexp.Compile
			}
			_, err := check(expr)
			if err != nil {
				if se, ok := err.(*syntax.Error); ok {
					pass.Reportf(val.Pos(), "%s in %q", se.Code, se.Expr)
				} else {
					pass.Reportf(val.Pos(), "%s", err)
				}
			}

			return true
		})
	}
	return nil, nil
}

func nodeFuncName(node *ast.CallExpr) string {
	se, ok := node.Fun.(*ast.SelectorExpr)
	if !ok {
		return ""
	}
	ie, ok := se.X.(*ast.Ident)
	if !ok {
		return ""
	}

	return fmt.Sprintf("%s.%s", ie.Name, se.Sel.Name)
}
