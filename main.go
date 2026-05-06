// beads - A package manager and build tool for Go projects
// Fork of gastownhall/beads with extended formula support
package main

import (
	"fmt"
	"os"

	"github.com/beads/beads/cmd"
)

const (
	version = "0.1.0"
	appName = "beads"
)

func main() {
	if err := cmd.Execute(version, appName); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
