package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/bodgit/sevenzip"
)

func main() {
	archivePath := "gibberish.7z"
	destDir := "/Users/osangote/Desktop/deep"

	r, err := sevenzip.OpenReader(archivePath)
	if err != nil {
		fmt.Println("Failed to open archive:", err)
		return
	}
	defer r.Close()

	// Try opening a file to check if it's encrypted
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			fmt.Printf("Archive %s is password protected or unreadable, skipping extraction.\n", archivePath)
			return
		}
		rc.Close()
		break // Successfully opened one file, assume not encrypted
	}

	// If not password protected, proceed to extract
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			fmt.Println("Failed to open file in archive:", err)
			continue
		}

		outPath := filepath.Join(destDir, f.Name)
		if err := os.MkdirAll(filepath.Dir(outPath), os.ModePerm); err != nil {
			fmt.Println("Failed to create directories:", err)
			rc.Close()
			continue
		}

		outFile, err := os.Create(outPath)
		if err != nil {
			fmt.Println("Failed to create output file:", err)
			rc.Close()
			continue
		}

		_, err = io.Copy(outFile, rc)
		if err != nil {
			fmt.Println("Failed to extract file:", err)
		}
		outFile.Close()
		rc.Close()
		fmt.Printf("Extracted: %s\n", outPath)
	}
}
