package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// :show start
func main() {
	nShown := 0
	err := filepath.Walk(".", func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if nShown > 4 {
			return nil
		}
		nShown++
		fmt.Printf("Path: %s, is dir: %v, size: %d bytes\n", fi.Name(), fi.IsDir(), fi.Size())
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk failed with '%s'\n", err)
	}
}

// :show end
