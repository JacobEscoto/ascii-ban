// Package style provides functions to pad ASCII string to the left, right, or center. It also provides functions to colorize ASCII text.
package style

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// renderAligned is a private helper that handles the terminal sizing and lipgloss styling.
func renderAligned(asciiArt string, fd int, alignment lipgloss.Position, fullscreen bool) string {
	width, height, err := term.GetSize(fd)
	if err != nil {
		return asciiArt // return original text if can't get the terminal size.
	}

	style := lipgloss.NewStyle().
		Width(width).
		Align(alignment)

	if fullscreen {
		style = style.
			Height(height).
			AlignVertical(lipgloss.Center)
	}

	cleanArt := strings.TrimRight(asciiArt, "\n")

	return style.Render(cleanArt)
}

// Left aligns the text to the left side of the terminal.
func Left(asciiArt string, fd int) string {
	return renderAligned(asciiArt, fd, lipgloss.Left, false)
}

// Center aligns the text to the center of the terminal.
func Center(asciiArt string, fd int) string {
	return renderAligned(asciiArt, fd, lipgloss.Center, false)
}

// Right aligns the text to the right side of the terminal.
func Right(asciiArt string, fd int) string {
	return renderAligned(asciiArt, fd, lipgloss.Right, false)
}

// FullscreenCenter is specifically designed for the live clock.
func FullscreenCenter(asciiArt string, fd int) string {
	return renderAligned(asciiArt, fd, lipgloss.Center, true)
}
