// Package cmd contains the complete definition of the CLI using the Cobra library.
//
// Provides the Execute() function as the entry point to `main.go`.
package cmd

import (
	"fmt"
	"strings"

	"github.com/JacobEscoto/ascii-ban/internal/font"
	"github.com/JacobEscoto/ascii-ban/internal/generator"
	"github.com/spf13/cobra"
)

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
		stdFont, fontErr := font.GetFont("standard")
		if fontErr != nil {
			return fontErr
		}
		result, err := generator.Render(usrText, stdFont)
		if err != nil {
			return err
		}
		fmt.Println(result)
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}
