package main

import (
	"log"
	"os"
	"time"
)

var logger *log.Logger

func init() {
	f, err := os.OpenFile("./inf.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		log.Panicln(err)
	}
	logger = log.Default()
	logger.SetOutput(f)

	logger.Println("logger ready.")
	log.Println("logger ready.")
}
func main() {
	for {
		// logger.Println("logger added.")
		logger.Println("???.")

		log.Println("logger added.")
		time.Sleep(time.Second)
	}
}
