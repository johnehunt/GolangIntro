package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// func main() {
// 	log.Println("Hello world!")
// 	log.Printf("A message for %s.", "John")
// }

func main() {
	var filename string = "logfile.log"
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		Formatter := new(log.TextFormatter)
		// You can change the Timestamp format. But you have to use the same date and time.
		// "2006-02-02 15:04:06" Works. If you change any digit, it won't work
		// ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
		Formatter.TimestampFormat = "02-01-2006 15:04:05"
		Formatter.FullTimestamp = true
		log.SetFormatter(Formatter)
		log.SetOutput(f)
		log.Info("Some information")
		log.Warning("This is a warning")
		log.Error("Not fatal. An error. Won't stop execution")
		log.Fatal("MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
		log.Panic("Do not panic")
	}
}
