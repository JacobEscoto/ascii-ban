package style

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Color returns the string with the specified color code.
func Color(asciiArt string, colorCode string) string {
	colorCode = formatColor(colorCode)
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(colorCode))
	return style.Render(asciiArt)
}

// isHex verifies that the color code contains only valid hexademical characters (0-9, a-f, A-F)
func isHex(colorCode string) bool {
	if len(colorCode) != 3 && len(colorCode) != 6 {
		return false
	}
	for _, char := range colorCode {
		isValidHex := (char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')

		if !isValidHex {
			return false
		}
	}
	return true
}

// formatColor ensures that the color code contains the '#' symbol if it is a hexademical code; otherwise, it returns the text exactly as it was received.
func formatColor(code string) string {
	code = strings.TrimSpace(code)

	if strings.HasPrefix(code, "#") {
		return code
	}

	if isHex(code) {
		return "#" + code
	}

	return code
}
