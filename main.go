package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var fileCategories = map[string][]string{
	"Images":    {".jpg", ".jpeg", ".png", ".gif"},
	"Documents": {".pdf", ".doc", ".docx", ".txt"},
	"Videos":    {".mp4", ".avi", ".mkv"},
	"Audio":     {".mp3", ".wav"},
	"Archives":  {".zip", ".tar", ".gz", ".rar", ".targz"},
}

func getFileCategory(extension string) string {
	for category, extensions := range fileCategories {
		for _, ext := range extensions {
			if strings.EqualFold(extension, ext) {
				return category
			}
		}
	}
	return "Others"
}

func organizeFiles(directory string) error {
	files, err := os.ReadDir(directory)
	if err != nil {
		return fmt.Errorf("failed to read directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			// Come back to allow recursive sorting/cleanup
			continue
		}

		extension := strings.ToLower(filepath.Ext(file.Name()))
		category := getFileCategory(extension)

		subDir := filepath.Join(directory, category)
		if _, err := os.Stat(subDir); os.IsNotExist(err) {
			err := os.Mkdir(subDir, os.ModePerm)
			if err != nil {
				log.Printf("Failed to create directory %s: %v", subDir, err)
				continue
			}
		}

		oldPath := filepath.Join(directory, file.Name())
		newPath := filepath.Join(subDir, file.Name())
		err := os.Rename(oldPath, newPath)
		if err != nil {
			log.Printf("Failed to move file %s: %v", file.Name(), err)
		} else {
			fmt.Printf("Moved %s to %s\n", file.Name(), category)
		}
	}

	return nil
}

func main() {
	// TODO: Replace with CL Args
	directory := "./test-files"

	fmt.Printf("Organizing files in directory: %s\n", directory)
	err := organizeFiles(directory)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("File organization complete!")
}
