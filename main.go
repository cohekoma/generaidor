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

	truncateStaticDir(publicRoot)

	err := filepath.WalkDir(contentRoot, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path != contentRoot {
			srcRelativePath, _ := filepath.Rel(contentRoot, path)                                     // -> remove `content` from the original path.
			destCompletePath := setGeneratedFileExtension(filepath.Join(publicRoot, srcRelativePath)) // -> constructed from relpath above, join the public with the rel path.
			destDirPath := filepath.Dir(destCompletePath)
			dirCreationErr := os.MkdirAll(destDirPath, 0755)
			generatedContent := generateStaticContent(path)

			fmt.Println(generatedContent)

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

// V1: Remove all child-dir and files inside the static (public) directory by delete the directory itself completely.
// @TODO: need to find a better way to handle this instead of removing everything then regenerate.
func truncateStaticDir(desiredDir string) {
	err := os.RemoveAll(desiredDir)

	if err != nil {
		log.Fatal("There seems to be an error while truncating the directory. Make sure you're not opening the directory or any file inside it.")
	}
}

// Take in a path to the original Markdown file.
// Generate the static output content to write into an HTML file.
func generateStaticContent(srcFilePath string) (outputContent string) {
	output, _ := os.ReadFile(srcFilePath)

	// @TODO:
	// Parse the frontmatter to html
	// Parse the Markdown format content to html
	// Map to the layout (if there's one)
	// Ouput the frontmatter and the html content to the html file.

	return string(output)
}
