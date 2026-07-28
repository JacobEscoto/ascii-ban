// Package align provides functions to pad ASCII strings to the left, right, or center.
package align

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// renderAligned is a private helper that handles the terminal sizing and lipgloss styling.
func renderAligned(asciiArt string, fd int, alignment lipgloss.Position, fullScreen bool) string {
	width, height, err := term.GetSize(fd)
	if err != nil {
		return asciiArt // return original text if it cannot get the terminal size.
	}

	style := lipgloss.NewStyle().
		Width(width).
		Align(alignment)

	if fullScreen {
		style = style.Height(height).AlignVertical(lipgloss.Center)
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

// FullScreenCenter is specifically designed for the live clock.
func FullScreenCenter(asciiArt string, fd int) string {
	return renderAligned(asciiArt, fd, lipgloss.Center, true)
}
