package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/JacobEscoto/ascii-ban/internal/align"
	"github.com/JacobEscoto/ascii-ban/internal/font"
	"github.com/JacobEscoto/ascii-ban/internal/generator"
	"github.com/spf13/cobra"
)

var clockCmd = &cobra.Command{
	Use:   "clock",
	Short: "Displays the time in real using ASCII numbers",
	Long:  `Displays a live clock directly in your terminal using ASCII art.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		font, fontErr := font.GetFont(generalOpts.font)
		if fontErr != nil {
			return fontErr
		}

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		fmt.Printf("\033[H\033[2J")

		defer fmt.Printf("\033[?25h")
		fd := int(os.Stdout.Fd())

		for range ticker.C {
			hour := getLocalTime()
			hourResult, err := generator.Render(hour, font)
			if err != nil {
				return fmt.Errorf("an error occurred while rendering the clock: %w", err)
			}

			centeredClock := align.FullScreenCenter(hourResult, fd)

			fmt.Printf("\033[H")
			fmt.Printf("%s", centeredClock)
		}
		return nil
	},
}

func getLocalTime() string {
	now := time.Now()
	return fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute())
}

func init() {
	rootCmd.AddCommand(clockCmd)
}
