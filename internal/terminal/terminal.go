package terminal

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func CenterText(asciiArt string, fd int) string {
	width, height, err := term.GetSize(fd)
	if err != nil {
		return asciiArt
	}

	termBox := lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center)

	cleanArt := strings.TrimRight(asciiArt, "\n")

	return termBox.Render(cleanArt)
}
