package importroll

import (
	"go/ast"
	"go/token"
	"go/types"
	"io/ioutil"
	"sync"

	"github.com/gobwas/glob"
	"golang.org/x/tools/go/analysis"
	"gopkg.in/yaml.v2"
)

var rule string

type Rule struct {
	Deny []string
}

var Analyzer = &analysis.Analyzer{
	Name:     "importroll",
	Doc:      "importroll",
	Run:      run,
	Requires: []*analysis.Analyzer{},
}

func init() {
	Analyzer.Flags.StringVar(&rule, "rule", "importroll.yml", "rule")
}

var lock sync.Mutex
var rules map[string]Rule

var globLock sync.Mutex
var globCache map[string]glob.Glob = make(map[string]glob.Glob)

func readRules() {
	if rules != nil {
		return
	}
	lock.Lock()
	defer lock.Unlock()
	bs, err := ioutil.ReadFile(rule)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bs, &rules)
	if err != nil {
		panic(err)
	}
}

func compileAndGetGlob(pattern string) glob.Glob {
	if v, found := globCache[pattern]; found {
		return v
	}
	globLock.Lock()
	defer globLock.Unlock()
	globCache[pattern] = glob.MustCompile(pattern)
	return globCache[pattern]
}

func run(pass *analysis.Pass) (interface{}, error) {
	readRules()
	path := pass.Pkg.Path()
	imports := pass.Pkg.Imports()
	deny := collectDeny(path, imports)
	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.GenDecl); ok {
				for _, spec := range decl.Specs {
					switch decl.Tok {
					case token.IMPORT:
						pkg := imported(pass.TypesInfo, spec.(*ast.ImportSpec))
						if _, found := deny[pkg.Path()]; found {
							pass.Reportf(spec.Pos(), "cannot import this package")
						}
					}
				}
			}
		}
	}
	return nil, nil
}

func collectDeny(path string, imports []*types.Package) map[string]struct{} {
	deny := map[string]struct{}{}
	for key, rule := range rules {
		g := compileAndGetGlob(key)
		if !g.Match(path) {
			continue
		}
		for _, v := range imports {
			importedPath := v.Path()
			for _, d := range rule.Deny {
				g := compileAndGetGlob(d)
				if g.Match(importedPath) {
					deny[importedPath] = struct{}{}
				}
			}
		}
	}
	return deny
}

func imported(info *types.Info, spec *ast.ImportSpec) *types.Package {
	obj, ok := info.Implicits[spec]
	if !ok {
		obj = info.Defs[spec.Name] // renaming import
	}
	return obj.(*types.PkgName).Imported()
}
