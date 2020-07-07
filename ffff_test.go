package ffff

import (
	"os"
	"strings"
	"testing"
)

func TestFuzzyFindPath(t *testing.T) {
	path, err := FuzzyFindPath("mon")
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(strings.ToLower(path), "mono") {
		t.Error("could not find font 'mono'")
	}
	if _, err := os.Stat(path); err != nil {
		t.Error(err)
	}
}
