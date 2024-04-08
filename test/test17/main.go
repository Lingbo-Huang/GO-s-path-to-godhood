package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("./toys/test17/app.log")
	if err != nil {
		log.Fatal("Error creating log file:", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Println("This is a log message.")
	log.Printf("This is a formatted log message with values: %d, %s\n", 11, "hh")
	log.Fatal("Ohno, fatal error.")
}
