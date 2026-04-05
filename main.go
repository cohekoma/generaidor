package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {

	err := handleStaticGeneration()

	if err != nil {
		log.Fatal("Fail to generate static content, please try again.")
	}
}

func handleStaticGeneration() error {
	contentRoot := "content"
	publicRoot := "public"

	err := filepath.WalkDir(contentRoot, func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path != contentRoot {
			srcRelativePath, _ := filepath.Rel(contentRoot, path)                                     // -> remove `content` from the original path.
			destCompletePath := setGeneratedFileExtension(filepath.Join(publicRoot, srcRelativePath)) // -> constructed from relpath above, join the public with the rel path.
			destDirPath := filepath.Dir(destCompletePath)

			fmt.Println(destCompletePath)

			dirCreationErr := os.MkdirAll(destDirPath, 0755)

			if dirCreationErr != nil {
				log.Fatal("There's been an error while creating directories, please check again.")
			}

		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// Take in a file path (as string). Change the extention fragment in the file path string from .md to .html
func setGeneratedFileExtension(orgFilePath string) (newFilePath string) {
	originalExtension := filepath.Ext(orgFilePath)
	staticExtension := ".html"

	newFilePath = orgFilePath[:len(orgFilePath)-len(originalExtension)] + staticExtension
	return newFilePath
}
