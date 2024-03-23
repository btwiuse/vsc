package main

import (
	"flag"
	"log"
	"os"

	"github.com/btwiuse/vsc"
)

func main() {
	err := vsc.Run(os.Args[1:])
	if err == flag.ErrHelp {
		return
	}
	if err != nil {
		log.Fatalln(err)
	}
}
