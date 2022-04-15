package main

import (
	"log"
	"os"

	"github.com/JonaVDM/wake-the-wizzard/wol"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No mac address provided")
	}

	mac := os.Args[1]
	wol.SendWol(mac)
}
