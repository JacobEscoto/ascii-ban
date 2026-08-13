// Package cmd contains the complete definition of the CLI using the Cobra library.
//
// Provides the Execute() function as the entry point to `main.go`.
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/JacobEscoto/ascii-ban/internal/font"
	"github.com/JacobEscoto/ascii-ban/internal/generator"
	"github.com/JacobEscoto/ascii-ban/internal/storage"
	"github.com/JacobEscoto/ascii-ban/internal/style"
	"github.com/spf13/cobra"
)

type generalOptions struct {
	outputPath string
	font       string
	align      string
}

var generalOpts = &generalOptions{}

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
		stdFont, fontErr := font.GetFont(generalOpts.font)
		if fontErr != nil {
			return fontErr
		}
		result, err := generator.Render(usrText, stdFont)
		if err != nil {
			return err
		}

		fd := int(os.Stdout.Fd())

		switch strings.ToLower(generalOpts.align) {
		case "center":
			result = style.Center(result, fd)

		case "right":
			result = style.Right(result, fd)

		case "left":
			result = style.Left(result, fd)

		case "fullscreen":
			result = style.FullscreenCenter(result, fd)
		default:
			return fmt.Errorf("invalid alignment '%s': options available are left, center, right", generalOpts.align)
		}

		if generalOpts.outputPath != "" {
			err := storage.WriteBanner(generalOpts.outputPath, result)
			if err != nil {
				return err
			}
			fmt.Printf("Banner successfully saved to %s\n", generalOpts.outputPath)
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
	rootCmd.Flags().StringVarP(&generalOpts.outputPath, "output", "o", "", "Generate an ASCII art banner in a text file")
	rootCmd.PersistentFlags().StringVarP(&generalOpts.font, "font", "f", "standard", "Specify the font to use")
	rootCmd.Flags().StringVar(&generalOpts.align, "align", "left", "Options available: left | right | center | fullscreen")
}
