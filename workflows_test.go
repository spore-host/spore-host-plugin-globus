package plugin_test

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// Every action CI runs is pinned to a full commit SHA with its version in a
// trailing comment. A floating tag (`actions/checkout@v4`) lets whoever controls
// the tag change what our CI executes; a SHA cannot move. The comment is what
// Dependabot rewrites alongside the SHA, and what lets a human read the pin.
var (
	usesLine  = regexp.MustCompile(`^\s*(?:-\s+)?uses:\s*(.*)$`)
	pinnedRef = regexp.MustCompile(`^[\w.-]+/[\w./-]+@[0-9a-f]{40} # v\d+\.\d+\.\d+$`)
)

func workflowFiles(t *testing.T) []string {
	t.Helper()
	var files []string
	for _, pat := range []string{"*.yml", "*.yaml"} {
		m, err := filepath.Glob(filepath.Join(".github", "workflows", pat))
		if err != nil {
			t.Fatal(err)
		}
		files = append(files, m...)
	}
	if len(files) == 0 {
		t.Fatal("no workflow files found under .github/workflows; a gate that checks nothing must not pass")
	}
	return files
}

func TestActionsArePinnedToSHAs(t *testing.T) {
	n := 0
	for _, f := range workflowFiles(t) {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		for i, l := range strings.Split(string(b), "\n") {
			m := usesLine.FindStringSubmatch(l)
			if m == nil {
				continue
			}
			n++
			ref := strings.TrimSpace(m[1])
			if strings.HasPrefix(ref, "./") || strings.HasPrefix(ref, "docker://") {
				continue // local action or image, not a mutable remote tag
			}
			if !pinnedRef.MatchString(ref) {
				t.Errorf("%s:%d: %q is not pinned as owner/repo@<40-hex SHA> # vX.Y.Z", f, i+1, ref)
			}
		}
	}
	if n == 0 {
		t.Fatal("found no `uses:` lines in any workflow; the pattern or the layout changed")
	}
}

// Pinning is only safe when something bumps the pins.
func TestDependabotCoversActions(t *testing.T) {
	b, err := os.ReadFile(filepath.Join(".github", "dependabot.yml"))
	if err != nil {
		t.Fatalf("read dependabot.yml: %v (actions are SHA-pinned; without Dependabot they go stale)", err)
	}
	if !regexp.MustCompile(`(?m)^\s*-\s*package-ecosystem:\s*["']?github-actions["']?\s*$`).Match(b) {
		t.Fatal("dependabot.yml has no github-actions entry, so the SHA-pinned actions are never bumped")
	}
}
