/* Package recheck provides a static checker for regular expressions

`receck` examines called to `regexp.*Compile*` and if the regular expression is
a literal string - will check that it's valid.

`recheck` will also examine lines that have a comment in the following format:

	r.HandleFunc("/articles/{category}/{id:[0-9]+}", handler) // recheck:0

Then number after "recheck:" is the argument to check (first one in this case)

*/
package recheck

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"
	"regexp/syntax"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Analyzer a regular expression analyzer
var Analyzer = &analysis.Analyzer{
	Name: "recheck",
	Doc:  "Validate regular expressions",
	Run:  run,
}

var (
	regexpFuncs = map[string]bool{
		"regexp.MustCompile":      true,
		"regexp.Compile":          true,
		"regexp.MustCompilePOSIX": true,
		"regexp.CompilePOSIX":     true,
	}
	commentRe = regexp.MustCompile(`recheck:\d+`)
)

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		checkedLines := comments(pass.Fset, file.Comments)
		ast.Inspect(file, func(node ast.Node) bool {
			ce, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			var check func(string) (*regexp.Regexp, error)
			argNum := -1

			lnum := pass.Fset.Position(node.Pos()).Line
			if n, ok := checkedLines[lnum]; ok {
				check = regexp.Compile
				argNum = n
			} else {
				funcName := nodeFuncName(ce)
				if regexpFuncs[funcName] {
					argNum = 0
					if strings.HasSuffix(funcName, "POSIX") {
						check = regexp.CompilePOSIX
					} else {
						check = regexp.Compile
					}
				}
			}

			if check == nil {
				return true
			}

			if argNum >= len(ce.Args) {
				// TODO: warning
				return true
			}

			val, ok := ce.Args[argNum].(*ast.BasicLit)
			if !ok {
				return true
			}

			// Trim enclosing "" or ``
			expr := strconv.Quote(val.Value[1 : len(val.Value)-1])
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

func comments(fset *token.FileSet, cg []*ast.CommentGroup) map[int]int {
	m := make(map[int]int) // lnum -> arg
	for _, g := range cg {
		for _, c := range g.List {
			text := commentRe.FindString(c.Text)
			if text == "" {
				continue
			}

			argNum := commentArg(text)
			if argNum == -1 {
				// TODO: warning
				continue
			}

			lnum := fset.Position(c.Slash).Line
			m[lnum] = argNum
		}
	}

	return m
}

// "// recheck:1" -> 1
func commentArg(text string) int {
	i := strings.Index(text, ":")
	if i == -1 {
		return -1
	}

	n, err := strconv.Atoi(text[i+1:])
	if err != nil {
		return -1
	}

	return n
}
