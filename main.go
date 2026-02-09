package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	contentRoot := "content"
	publicRoot := "public"

	err := filepath.WalkDir(contentRoot, func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path != contentRoot {
			if !dir.IsDir() {
				// err := os.Mkdir(publicRoot+dir.Name(), 0755)
				srcFile := dir.Name()
				srcFileExt := filepath.Ext(srcFile)
				destFile := srcFile[:len(srcFile)-len(srcFileExt)] + ".html"

				generatedFile, err := os.Create(publicRoot + "/" + destFile)

				if err != nil {
					panic(err)
				}

				fmt.Println("Generated file: " + generatedFile.Name())
				// fmt.Println("Making file in: " + publicRoot + dir.Name())
			}
		}

		// fmt.Println(path)
		return nil
	})

	if err != nil {
		panic(err)
	}
}
