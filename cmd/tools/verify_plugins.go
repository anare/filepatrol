package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var pluginPaths = []string{
	"auth/jwt",
	"poster/http",
	"processor/simple",
	"watcher/local_folder",
}

func main() {
	base := "plugins"

	errors := 0
	for _, plugin := range pluginPaths {
		parts := strings.Split(plugin, "/")
		if len(parts) < 2 {
			fmt.Printf("Invalid plugin path format: %s\n", plugin)
			errors++
			continue
		}
		pluginName := parts[len(parts)-1]
		expected := filepath.Join(base, plugin, "main", pluginName+".go")

		if _, err := os.Stat(expected); os.IsNotExist(err) {
			fmt.Printf("❌ Missing file: %s\n", expected)
			errors++
		} else {
			fmt.Printf("✅ Found: %s\n", expected)
		}
	}

	if errors > 0 {
		fmt.Printf("\n%d issues found.\n", errors)
		os.Exit(1)
	} else {
		fmt.Println("\nAll plugin structures look correct.")
	}
}
