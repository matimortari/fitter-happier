package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
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
	"coverage":     true,
	"tmp":          true,
	"temp":         true,
	"log":          true,
	"logs":         true,
	"test-results": true,
}

func main() {
	excludeFlag := flag.String("exclude", "", "Comma-separated names to exclude")
	noDefaultExclude := flag.Bool("no-default-exclude", false, "Disable default exclusions")
	flag.Parse()

	path := "."
	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}

	exclude := make(map[string]bool)
	if !*noDefaultExclude {
		for name := range defaultExclude {
			exclude[name] = true
		}
	}

	if *excludeFlag != "" {
		for _, name := range strings.Split(*excludeFlag, ",") {
			name = strings.TrimSpace(name)
			if name != "" {
				exclude[name] = true
			}
		}
	}

	if _, err := os.Stat(path); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	printTree(path, "", exclude)
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
