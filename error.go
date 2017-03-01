package jata

import (
	"log"
	"os"
)

var (
	// Error is the stderr logger
	Error *log.Logger
)

func init() {
	Error = log.New(os.Stderr,
		"jata: ",
		log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)
}
