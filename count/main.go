package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
)

func main() {

	// Process all *.go files in current directory.
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("directory: %v\n", dir)

	fset := token.NewFileSet()
	//f, err := parser.ParseFile(fset, "src.go", src, 0)
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		fmt.Printf("  package: %v\n", pkg.Name)
		for fileName, file := range pkg.Files {
			fmt.Printf("    file: %v\n", fileName)

			// "Properties" (props) of and identifier are:
			// count, position of first and last occurrence (span).
			type idProps struct {
				Count    int
				FirstPos token.Pos
				LastPos  token.Pos
			}

			// We need both a map and a slice of properties.
			var idCountersMap map[string]idProps

			type idNameAndProps struct {
				Name string
				idProps
			}
			var idCountsSlice []idNameAndProps

			// Initialise counters.
			idCountersMap = make(map[string]idProps)

			// Inspect the AST and count identifiers.
			ast.Inspect(file, func(n ast.Node) bool {
				var s string
				// Count identifiers.
				switch x := n.(type) {
				// case *ast.BasicLit:
				//	s = x.Value
				case *ast.Ident:
					s = x.Name
					p := idCountersMap[s]
					p.Count = p.Count + 1
					if p.FirstPos == 0 {
						p.FirstPos = n.Pos()
					} else if n.Pos() < p.FirstPos {
						p.FirstPos = n.Pos()
					}
					if n.Pos() > p.LastPos {
						p.LastPos = n.Pos()
					}
					idCountersMap[s] = p
				}
				return true
			})

			// Create a slice from map.
			for k, v := range idCountersMap {
				idCountsSlice = append(idCountsSlice, idNameAndProps{k, v})
			}

			// Sort the slice.
			sort.Slice(idCountsSlice,
				func(i, j int) bool {
					return idCountsSlice[i].Count > idCountsSlice[j].Count
				},
			)

			for _, kv := range idCountsSlice {
				fmt.Printf("      %v : %v (%v)\n",
					kv.Name, kv.Count, kv.LastPos-kv.FirstPos+1)
			}

		}
	}

}

// Notes:
// How to sort a map?
// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values

// Instrumenting Go code via AST
// https://developers.mattermost.com/blog/instrumenting-go-code-via-ast/
