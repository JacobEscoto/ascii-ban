// Package generator implements horizontal FIGlet font rendering.
//
// Usage:
//
// result, err := generator.Render("Your Text", font.Font)
package generator

import (
	"errors"
	"strings"

	"github.com/JacobEscoto/ascii-ban/internal/font"
)

// Render takes the text and renders it by iterating and building the final result.
// If a character does not exist, it receives an alternative character, usually '?' or a space.
func Render(text string, chosenFont font.Font) (string, error) {
	var builder strings.Builder

	if text == "" {
		return "", errors.New("the text to be rendered is empty")
	}

	if chosenFont.UppercaseOnly {
		text = strings.ToUpper(text)
	}

	fontCharacters := chosenFont.Characters

	for idx := range chosenFont.Height {
		for _, character := range text {
			asciiChar, exists := fontCharacters[character]
			if !exists {
				if asciiChar, exists = fontCharacters['?']; !exists {
					asciiChar = fontCharacters[' ']
				}
			}

			lnChar := asciiChar[idx]

			builder.WriteString(lnChar)
		}
		builder.WriteString("\n")
	}

	return builder.String(), nil
}
