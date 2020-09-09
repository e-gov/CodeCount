package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
)

func main() {

	// Parse command line flags.
	var cFlag = flag.Bool("c", false, "-c - Output only comments")
	var hFlag = flag.Bool("h", false, "-c - Output only comments")
	flag.Parse()

	if *hFlag {
		fmt.Println(`Usage:
		-c - Output only comments`)
		os.Exit(0)
	}

	// Get current directory.
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("directory: %v\n", dir)

	// Parse all *.go files in the current directory.
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Walk all parsed packages.
	for _, pkg := range pkgs {
		fmt.Printf("  package: %v\n", pkg.Name)
		// Walk all files of the package.
		for fileName, file := range pkg.Files {
			fmt.Printf("    file: %v\n", fileName)

			// "Properties" (props) of and identifier:
			type idProps struct {
				// Count of occurrencies.
				Count int
				// Position of first and last occurrence (span).
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
				case *ast.Ident:
					if !*cFlag {
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

			// Print the slice (identifier counts and spans).
			for _, kv := range idCountsSlice {
				fmt.Printf("      %v : %v (%v)\n",
					kv.Name, kv.Count, kv.LastPos-kv.FirstPos+1)
			}

			// Print file's comments.
			fmt.Println("COMMENTS:")
			for _, cGroup := range file.Comments {
				for _, c := range cGroup.List {
					fmt.Printf("%v\n", c.Text)
				}

			}

		}
	}

}

// Notes:
// How to sort a map?
// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values

// Instrumenting Go code via AST
// https://developers.mattermost.com/blog/instrumenting-go-code-via-ast/
