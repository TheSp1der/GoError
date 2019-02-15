package goerror

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// Fatal reports error on stdout and halts the program
func Fatal(err error) {
	var (
		Error *log.Logger // error handler
	)

	Error = log.New(
		os.Stderr,
		color.RedString("ERROR")+":\t",
		log.Ldate|log.Ltime,
	)

	Error.Println(err)
	os.Exit(1)
}

// FatalReceiver reports error on stdout and halts the program via a blocking channel
func FatalReceiver(err <-chan error) {
	for {
		Fatal(<-err)
	}
}
