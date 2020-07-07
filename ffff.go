package ffff

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/beta/freetype/truetype"
	"github.com/k1LoW/fontdir"
	"github.com/sahilm/fuzzy"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

type Font struct {
	path string
	face font.Face
}

// FuzzyFind find font by keyword
func FuzzyFind(keyword string, to *truetype.Options, oo *opentype.FaceOptions) (Font, error) {
	list := []string{}
	names := []string{}
	paths := []string{}
	fonts := map[string]Font{}
	pathOnly := false
	lk := strings.ToLower(keyword)
	if strings.HasSuffix(lk, ".ttf") || strings.HasSuffix(lk, ".otf") {
		pathOnly = true
	}
	for _, dir := range fontdir.Get() {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			lp := strings.ToLower(path)
			abs, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			var face font.Face
			if strings.HasSuffix(lp, ".ttf") {
				// TrueType
				d, err := ioutil.ReadFile(filepath.Clean(path))
				if err != nil {
					return err
				}
				f, err := truetype.Parse(d)
				if err != nil {
					return nil
				}
				name := f.Name(4)
				names = append(names, name)
				face = truetype.NewFace(f, to)
				if !pathOnly {
					fonts[name] = Font{
						path: abs,
						face: face,
					}
				}
			} else if strings.HasSuffix(lp, ".otf") {
				// OpenType
				d, err := ioutil.ReadFile(filepath.Clean(path))
				if err != nil {
					return err
				}
				f, err := sfnt.Parse(d)
				if err != nil {
					return nil
				}
				name, err := f.Name(nil, 4)
				if err != nil {
					return nil
				}
				names = append(names, name)
				face, err = opentype.NewFace(f, oo)
				if err != nil {
					return nil
				}
				if !pathOnly {
					fonts[name] = Font{
						path: abs,
						face: face,
					}
				}
			} else {
				return nil
			}
			paths = append(paths, abs)
			fonts[abs] = Font{
				path: abs,
				face: face,
			}
			return nil
		})
		if err != nil {
			return Font{}, err
		}
	}

	list = append(list, names...)
	list = append(list, paths...)

	matches := fuzzy.Find(keyword, list)
	if len(matches) == 0 {
		return Font{}, fmt.Errorf("could not find font: %s", keyword)
	}
	m, ok := fonts[matches[0].Str]
	if !ok {
		return Font{}, fmt.Errorf("could not find font: %s", keyword)
	}
	return m, nil
}

// FuzzyFindPath find font file path by keyword
func FuzzyFindPath(keyword string) (string, error) {
	f, err := FuzzyFind(keyword, nil, nil)
	if err != nil {
		return "", err
	}
	return f.path, nil
}

// FuzzyFindFace find font.Face by keyword
func FuzzyFindFace(keyword string, to *truetype.Options, oo *opentype.FaceOptions) (font.Face, error) {
	f, err := FuzzyFind(keyword, nil, nil)
	if err != nil {
		return &opentype.Face{}, err
	}
	return f.face, nil
}
