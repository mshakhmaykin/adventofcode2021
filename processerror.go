package main

import (
	"log"
	"os"
)

func processerror(err *error) {
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
