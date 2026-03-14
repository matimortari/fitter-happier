package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

var defaultExclude = map[string]bool{
	"node_modules": true,
	"dist":         true,
	".git":         true,
	".nuxt":        true,
	".next":        true,
	".output":      true,
	".venv":        true,
	"out":          true,
	"build":        true,
	"bin":          true,
	"obj":          true,
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	printTree(path, "", defaultExclude)
}

func printTree(basePath, indent string, exclude map[string]bool) {
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return
	}

	var filtered []os.DirEntry
	for _, entry := range entries {
		if exclude[entry.Name()] {
			continue
		}

		filtered = append(filtered, entry)
	}

	sort.Slice(filtered, func(i, j int) bool {
		if filtered[i].IsDir() != filtered[j].IsDir() {
			return filtered[i].IsDir()
		}
		return filtered[i].Name() < filtered[j].Name()
	})

	for i, entry := range filtered {
		isLast := i == len(filtered)-1

		connector := "├── "
		newIndent := indent + "│   "
		if isLast {
			connector = "└── "
			newIndent = indent + "    "
		}

		fmt.Printf("%s%s%s\n", indent, connector, entry.Name())

		if entry.IsDir() {
			printTree(filepath.Join(basePath, entry.Name()), newIndent, exclude)
		}
	}
}
