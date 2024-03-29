package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	timeNow := time.Now().Add(response.ClockOffset)

	fmt.Printf("Ntp: %s\n", timeNow.Local().Format("2006.01.02 15:04:05"))
	fmt.Printf("Current time: %s\n", time.Now().Format("2006.01.02 15:04:05"))
	os.Exit(1)
}
