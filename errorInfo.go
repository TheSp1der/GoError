package goerror

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// Info reports info on stdout via a blocking channel
func Info(err error) {
	var (
		Info *log.Logger // warning handler
	)

	Info = log.New(
		os.Stdout,
		color.GreenString("INFO")+":\t",
		log.Ldate|log.Ltime,
	)

	Info.Println(err)
}

// InfoReceiver reports info on stdout via a blocking channel
func InfoReceiver(err <-chan error) {
	for {
		Fatal(<-err)
	}
}
