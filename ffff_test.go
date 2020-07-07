package ffff

import (
	"os"
	"strings"
	"testing"
)

func TestFuzzyFindPath(t *testing.T) {
	path, err := FuzzyFindPath("gothi")
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(strings.ToLower(path), "gothic") {
		t.Error("could not find font 'gothic'")
	}
	if _, err := os.Stat(path); err != nil {
		t.Error(err)
	}
}
