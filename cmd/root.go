// Package cmd contains the complete definition of the CLI using the Cobra library.
//
// Provides the Execute() function as the entry point to `main.go`.
package cmd

import (
	"fmt"
	"strings"

	"github.com/JacobEscoto/ascii-ban/internal/font"
	"github.com/JacobEscoto/ascii-ban/internal/generator"
	"github.com/JacobEscoto/ascii-ban/internal/storage"
	"github.com/spf13/cobra"
)

type Options struct {
	outputPath string
	font       string
}

var options = &Options{}

var rootCmd = &cobra.Command{
	Use:   "ascii-ban <text>",
	Short: "Generate beautiful ASCII art banners directly in your terminal.",
	Long: `ascii-ban is a lightweight command-line tool that transforms standard text
into stylized ASCII art banners.

Perfect for creating custom script headers, or just having fun in the command line. 
It automatically handles spaces, joins multiple arguments, and uses smart fallback 
characters if a symbol is not supported by the chosen font.

Example:
  ascii-ban Hello World`,

	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		usrText := strings.Join(args, " ")
		stdFont, fontErr := font.GetFont(options.font)
		if fontErr != nil {
			return fontErr
		}
		result, err := generator.Render(usrText, stdFont)
		if err != nil {
			return err
		}
		if options.outputPath != "" {
			err := storage.WriteBanner(options.outputPath, result)
			if err != nil {
				return err
			}
			fmt.Printf("Banner successfully saved to %s\n", options.outputPath)
			return nil
		}

		fmt.Println(result)
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&options.outputPath, "output", "o", "", "Generate an ASCII art banner in a text file")
	rootCmd.Flags().StringVarP(&options.font, "font", "f", "standard", "Choose a font to display the ASCII banner. standard | slant")
}
