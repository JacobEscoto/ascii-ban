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
	color      string
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

Examples:
  ascii-ban Hello World
  ascii-ban Hello World -f ansi-shadow
  ascii-ban Hello World -f ansi-shadow -c "#1DF232"`,

	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		usrText := strings.Join(args, " ")
		chosenFont, fontErr := font.GetFont(generalOpts.font)
		if fontErr != nil {
			return fontErr
		}
		renderedASCII, err := generator.Render(usrText, chosenFont)
		if err != nil {
			return err
		}

		fd := int(os.Stdout.Fd())

		switch strings.ToLower(generalOpts.align) {
		case "center":
			renderedASCII = style.Center(renderedASCII, fd)

		case "right":
			renderedASCII = style.Right(renderedASCII, fd)

		case "left":
			renderedASCII = style.Left(renderedASCII, fd)

		case "fullscreen":
			renderedASCII = style.FullscreenCenter(renderedASCII, fd)
		default:
			return fmt.Errorf("invalid alignment '%s': options available are left, center, right", generalOpts.align)
		}

		renderedASCII = style.Color(renderedASCII, generalOpts.color)

		if generalOpts.outputPath != "" {
			err := storage.WriteBanner(generalOpts.outputPath, renderedASCII)
			if err != nil {
				return err
			}
			fmt.Printf("Banner successfully saved to %s\n", generalOpts.outputPath)
			return nil
		}

		fmt.Println(renderedASCII)
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
	rootCmd.PersistentFlags().StringVarP(&generalOpts.color, "color", "c", "white", "Specify the color to print the rendered text")
}
