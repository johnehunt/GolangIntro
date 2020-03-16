package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func checkForError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("Starting")

	// Create a file and write a simple string to it
	f, err := os.Create("test1.txt")
	checkForError(err)

	l, err := f.WriteString("Hello World")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l)

	fmt.Println("--------------")

	// Write array of bytes to a file
	d1 := []byte("hello Go World\n")
	err = ioutil.WriteFile("data1.txt", d1, 0644)
	checkForError(err)

	// Can also use a buffered writer

	f, err = os.Create("data2.txt")
	checkForError(err)
	w := bufio.NewWriter(f)
	numOfBytes, err := w.WriteString("buffered\n")
	checkForError(err)
	fmt.Printf("wrote %d bytes\n", numOfBytes)
	w.Flush()

	fmt.Println("Done")

}
