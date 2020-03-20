package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("Starting")

	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Get the extensions
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println("Done")
}
