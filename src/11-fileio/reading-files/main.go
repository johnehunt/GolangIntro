package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Read contents of file into memory
	data, err := ioutil.ReadFile("input.txt")
	PanicIfError(err)
	fmt.Println(string(data))

	fmt.Println("-------------")

	// For more control - read a line at a time
	f, err := os.Open("input.txt")
	PanicIfError(err)
	// Close the file when you’re done
	defer f.Close()

	bufferedReader := bufio.NewReader(f)
	for {
		line, err := bufferedReader.ReadString('\n')
		if err == io.EOF {
			fmt.Print("read: ", line)
			break
		} else {
			fmt.Print("read: ", line)
		}
	}

	fmt.Println("Done")

}
