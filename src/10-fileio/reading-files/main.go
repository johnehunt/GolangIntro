package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Read contents of file into memory
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	fmt.Println(string(dat))

	fmt.Println("-------------")

	// For more control - read a line at a time
	f, err := os.Open("input.txt")
	check(err)

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

	//	Close the file when you’re done
	f.Close()
	fmt.Println("Done")

}
