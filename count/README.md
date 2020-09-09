Program `CodeCount` implements some ideas of [3 Things That Counting Words Can Reveal on Your Code](https://www.fluentcpp.com/2018/10/09/3-things-that-counting-words-can-reveal-on-your-code/).

`CodeCount` can be operated in two modes:

`count` analyzes all *.go files in current directory and prints a sorted list of identifier counts. Identifier "span" (distance between first and last occurrence of the identifier in fileset) is also printed. Example output (excerpt):

````
C:\TÖÖS\CodeCount\count> go run .
directory: C:\TÖÖS\CodeCount\count
  package: main
    file: C:\TÖÖS\CodeCount\count\main.go
      p : 10 (282)
      idCountsSlice : 7 (1041)
      Pos : 7 (789)
      n : 7 (397)
      fmt : 7 (2324)
      Count : 6 (1408)
      err : 6 (277)
      FirstPos : 6 (1360)
      Printf : 5 (2156)
      idCountersMap : 5 (830)
      kv : 5 (105)
````

`count -c` prints comments of the fileset. Example:

````
C:\TÖÖS\CodeCount\count> go run . -c
directory: C:\TÖÖS\CodeCount\count
  package: main
    file: C:\TÖÖS\CodeCount\count\main.go
COMMENTS:
// Parse command line flags.
// Get current directory.
// Parse all *.go files in the current directory.
// Walk all parsed packages.
// Walk all files of the package.
// "Properties" (props) of and identifier:
// Count of occurrencies.
// Position of first and last occurrence (span).
// We need both a map and a slice of properties.
// Initialise counters.
// Inspect the AST and count identifiers.
// Count identifiers.
// Create a slice from map.
// Sort the slice.
// Print the slice (identifier counts and spans).
// Print file's comments.
// Notes:
// How to sort a map?
// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
// Instrumenting Go code via AST
// https://developers.mattermost.com/blog/instrumenting-go-code-via-ast/
````