package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	contentRoot := "content"
	// publicRoot := "public"

	err := filepath.WalkDir(contentRoot, func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if dir.IsDir() {
			return nil
		}

		fmt.Println(path)
		return nil
	})

	if err != nil {
		panic(err)
	}
}
