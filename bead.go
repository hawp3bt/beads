package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Bead represents a single installable formula unit.
// Each bead corresponds to a .formula.toml file in the .beads/formulas directory.
type Bead struct {
	Name    string
	Version string
	Formula string // path to the .formula.toml file
}

// BeadStore manages the local bead registry and formula resolution.
type BeadStore struct {
	FormulasDir string
	Beads       map[string]*Bead
}

// NewBeadStore initializes a BeadStore rooted at the given formulas directory.
func NewBeadStore(formulasDir string) (*BeadStore, error) {
	info, err := os.Stat(formulasDir)
	if err != nil {
		return nil, fmt.Errorf("formulas directory not found: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", formulasDir)
	}

	store := &BeadStore{
		FormulasDir: formulasDir,
		Beads:       make(map[string]*Bead),
	}
	if err := store.load(); err != nil {
		return nil, err
	}
	return store, nil
}

// load scans the formulas directory and registers all discovered formula files.
func (bs *BeadStore) load() error {
	entries, err := os.ReadDir(bs.FormulasDir)
	if err != nil {
		return fmt.Errorf("failed to read formulas dir: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(name, ".formula.toml") {
			continue
		}
		// Derive bead name by stripping the suffix
		beadName := strings.TrimSuffix(name, ".formula.toml")
		bs.Beads[beadName] = &Bead{
			Name:    beadName,
			Formula: filepath.Join(bs.FormulasDir, name),
		}
	}
	return nil
}

// Get returns the Bead with the given name, or an error if not found.
func (bs *BeadStore) Get(name string) (*Bead, error) {
	b, ok := bs.Beads[name]
	if !ok {
		return nil, fmt.Errorf("bead %q not found in store", name)
	}
	return b, nil
}

// List returns all bead names registered in the store, sorted alphabetically.
// Sorted output makes it easier to scan when you have a lot of formulas.
func (bs *BeadStore) List() []string {
	names := make([]string, 0, len(bs.Beads))
	for name := range bs.Beads {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
