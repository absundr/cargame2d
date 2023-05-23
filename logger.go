package main

import (
	"os"
	"time"
)

func Log(input string) {
	// Open the file for writing.
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
	  os.Exit(0)
	}
  
	t := time.Now()
  	timestamp := t.Format("2006-01-02 15:04:05")

	// Write the text to the file.
	_, err = f.WriteString(timestamp+": "+input + "\n")
	if err != nil {
		os.Exit(0)
	}
  
	// Close the file.
	err = f.Close()
	if err != nil {
		os.Exit(0)
	}
}