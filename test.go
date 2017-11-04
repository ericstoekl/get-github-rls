package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filePath := "./"
	filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
