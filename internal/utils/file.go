package utils

import (
	"fmt"
	"io"
	"os"
)

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	return nil
}

func WriteToFile(filePath string, reader io.Reader) error {
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer outFile.Close()

	if _, err = io.Copy(outFile, reader); err != nil {
		return fmt.Errorf("failed to write data to file %s: %w", filePath, err)
	}

	return nil
}
