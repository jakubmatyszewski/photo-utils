package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the directory path as an argument.")
	}

	// Get the directory from the command-line arguments
	dir := os.Args[1]

	// List files in the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	totalFiles := len(files)

	// Loop through the files
	for i, file := range files {
		fmt.Printf("Processing %d/%d\n", i+1, totalFiles)

		filename := file.Name()
		fileExt := filepath.Ext(filename)
		if !file.IsDir() && (strings.ToLower(fileExt) == ".jpg" || strings.ToLower(fileExt) == ".jpeg" || strings.ToLower(fileExt) == ".png") {
			// Get the full file path
			filePath := filepath.Join(dir, file.Name())

			fmt.Println("Processing file:", filePath)

			// Execute the `cjxl` command with the file path
			opts := strings.Split("-e 7 -v --lossless_jpeg=0 -q 90", " ")

			outputPath := strings.Replace(filePath, fileExt, ".JXL", -1)
			cmd := exec.Command("cjxl", append([]string{filePath, outputPath}, opts...)...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("Failed to execute command for %s: %s", filePath, err)
			}

			fmt.Printf("Command output for %s:\n%s\n", filePath, output)
		}
	}
}
