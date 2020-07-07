package ffff

import (
	"strings"
	"testing"
)

func TestFuzzyFindPath(t *testing.T) {
	path, _ := FuzzyFindPath("Ari")
	if !strings.Contains(strings.ToLower(path), "arial") {
		t.Error("could not find font 'arial'")
	}
}
