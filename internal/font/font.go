// Package font provides functions for adding new fonts to the tool
// and retrieving the established fonts.
package font

import "errors"

// Font represents a struct that includes a name and a characters map that returns an array of strings.
type Font struct {
	Name       string
	Characters map[rune][]string
	Height     int
}

var AvailableFonts = make(map[string]Font)

// AddFont adds a new font, checking that is not already added.
func AddFont(font Font) error {
	if _, exists := AvailableFonts[font.Name]; exists {
		return errors.New("font already exists")
	}

	AvailableFonts[font.Name] = font
	return nil
}

// GetFont searches for an available font in the AvailableFonts map.
// It validates that the name of the font is not empty and that it actually exists.
func GetFont(name string) (Font, error) {
	if name == "" {
		return Font{}, errors.New("font name not specified")
	}
	font, exists := AvailableFonts[name]
	if !exists {
		return Font{}, errors.New("the font specified is not available")
	}

	return font, nil
}
