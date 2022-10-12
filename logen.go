package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	logFile, _ := os.OpenFile("./out.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	log.SetOutput(logFile)

	i := 1

	for {
		log.Println("[ ERROR] bzzz " + strconv.Itoa(i))

		time.Sleep(1 * time.Millisecond)

		i++
	}
}
