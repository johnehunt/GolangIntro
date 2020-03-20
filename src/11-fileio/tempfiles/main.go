package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)


func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Create a temp dir in the system default temp folder
	tempDirPath, err := ioutil.TempDir("", "myTempDir")
	PanicIfError(err)
	fmt.Println("Temp dir created:", tempDirPath)

	// Create a file in new temp directory
	tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
	PanicIfError(err)
	fmt.Println("Temp file created:", tempFile.Name())

	// Close file
	err = tempFile.Close()
	PanicIfError(err)

	// Delete the resources we created
	err = os.Remove(tempFile.Name())
	PanicIfError(err)
	
	err = os.Remove(tempDirPath)
	PanicIfError(err)
}
