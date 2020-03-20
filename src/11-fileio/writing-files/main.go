package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("Starting")

	// Write array of bytes to a file
	d1 := []byte("hello Go World\n")
	err = ioutil.WriteFile("data1.txt", d1, 0644)
	PanicIfError(err)

	fmt.Println("--------------")

	// Create a file and write a simple string to it
	f, err := os.Create("test1.txt")
	PanicIfError(err)

	l, err := f.WriteString("Hello World")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l)

	// Can also use a buffered writer

	f, err = os.Create("data2.txt")
	PanicIfError(err)
	w := bufio.NewWriter(f)
	numOfBytes, err := w.WriteString("buffered\n")
	PanicIfError(err)
	fmt.Printf("wrote %d bytes\n", numOfBytes)
	w.Flush()

	fmt.Println("Done")

}
