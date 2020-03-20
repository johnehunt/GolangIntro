package main

import (
	"fmt"
	"os"
)

func PanicIfError(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {
	file, _ := os.Open("input.txt")

	var offset int64 = 2
	var whence int = 0

	_, err := file.Seek(offset, whence)
	PanicIfError(err)

	bytes := make([]byte, 4)
	n1, err := file.Read(bytes)
	PanicIfError(err)
	fmt.Printf("%d bytes: %s\n", n1, string(bytes[:n1]))

	offset2, err := file.Seek(10, 0)
	PanicIfError(err)

	b2 := make([]byte, 10)
	n2, err := file.Read(b2)
	PanicIfError(err)
	fmt.Printf("%d bytes @ %d: ", n2, offset2)
	fmt.Printf("%v\n", string(b2[:n2]))

	file.Close()

}
