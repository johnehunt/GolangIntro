package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Handle error situations
func PanicIfError(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {

	err := os.Mkdir("subdir", 0755)
	PanicIfError(err)

	defer os.RemoveAll("subdir")

	err = os.MkdirAll("subdir/parent/child", 0755)
	PanicIfError(err)

	// ReadDir lists directory contents
	contents, err := ioutil.ReadDir("..")
	PanicIfError(err)
	for _, entry := range contents {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	visitor := func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(" ", p, info.IsDir())
		return nil
	}
	err = filepath.Walk("..", visitor)

}
