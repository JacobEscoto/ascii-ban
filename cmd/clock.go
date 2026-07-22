package cmd

import (
	"fmt"
	"time"

	"github.com/JacobEscoto/ascii-ban/internal/font"
	"github.com/JacobEscoto/ascii-ban/internal/generator"
	"github.com/spf13/cobra"
)

type ClockOptions struct {
	font string
}

var clockOptions = ClockOptions{}

var clockCmd = &cobra.Command{
	Use:   "clock",
	Short: "Displays the time in real using ASCII numbers",
	Long:  `Displays a live clock directly in your terminal using ASCII art.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		font, fontErr := font.GetFont(clockOptions.font)
		if fontErr != nil {
			return fontErr
		}

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		fmt.Print("\033[H\033[2J")

		for range ticker.C {
			hour := getLocalTime()
			hourResult, err := generator.Render(hour, font)
			if err != nil {
				return fmt.Errorf("an error occurred while rendering the clock: %w", err)
			}
			fmt.Print("\033[H")
			fmt.Print(hourResult)
		}
		return nil
	},
}

func getLocalTime() string {
	now := time.Now()
	return fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute())
}

func init() {
	clockCmd.Flags().StringVarP(&clockOptions.font, "font", "f", "standard", "Available fonts: standard | slant")
	rootCmd.AddCommand(clockCmd)
}
