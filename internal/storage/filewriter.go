// Package storage provides functions for saving ASCII text to a specified path.
//
// Its primary responsibility is to receive content and a destination path,
// and persist that content in the file system.
// By default, the operation overwrites the file if it already exists (truncating it).
package storage

import (
	"fmt"
	"os"
)

// WriteBanner saves the ASCII text into a file at the specified path.
// If the file already exists, it will overwrite it.
func WriteBanner(filePath, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0o644)
	if err != nil {
		return fmt.Errorf("failed to write the banner to file: %w", err)
	}
	return nil
}
