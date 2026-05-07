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
		// Print error to stderr and exit with a non-zero status code.
		// Using exit code 2 to distinguish usage/runtime errors from
		// the default exit code 1 used by the OS for general errors.
		//
		// NOTE: exit code 1 is reserved for cases where the OS itself
		// signals failure; exit code 2 aligns with how tools like
		// grep and diff signal misuse or runtime errors.
		fmt.Fprintf(os.Stderr, "%s: error: %v\n", appName, err)
		os.Exit(2)
	}
}
