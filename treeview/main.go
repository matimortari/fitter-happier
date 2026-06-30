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
	// VCS
	".git": true,
	".hg":  true,
	".svn": true,

	// JS / TS
	"node_modules":     true,
	"bower_components": true,
	".next":            true,
	".nuxt":            true,
	".output":          true,
	".svelte-kit":      true,
	".angular":         true,
	".turbo":           true,
	".vercel":          true,
	".parcel-cache":    true,
	".yarn":            true,
	".pnpm-store":      true,

	// Python
	"__pycache__":   true,
	".venv":         true,
	"venv":          true,
	".tox":          true,
	".pytest_cache": true,
	".mypy_cache":   true,
	".ruff_cache":   true,
	".eggs":         true,
	"htmlcov":       true,
	"site-packages": true,

	// Ruby / PHP
	".bundle": true,
	"vendor":  true,

	// Java / Kotlin / Scala / Rust / C++
	"target":              true,
	".gradle":             true,
	".bloop":              true,
	".metals":             true,
	"cmake-build-debug":   true,
	"cmake-build-release": true,
	"CMakeFiles":          true,

	// .NET
	"bin":      true,
	"obj":      true,
	"packages": true,
	".vs":      true,

	// Elixir / Erlang
	"_build":     true,
	"deps":       true,
	".elixir_ls": true,

	// Haskell
	".stack-work":   true,
	"dist-newstyle": true,

	// Swift / Dart / Flutter
	".build":      true,
	"DerivedData": true,
	"Pods":        true,
	".swiftpm":    true,
	".dart_tool":  true,
	".pub-cache":  true,

	// Terraform / misc tooling
	".terraform": true,
	".direnv":    true,
	".cache":     true,
	".idea":      true,

	// Common build / output / logs
	"dist":         true,
	"out":          true,
	"build":        true,
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
