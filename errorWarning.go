package goerror

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// Warning reports warning on stdout
func Warning(err error) {
	var (
		Warning *log.Logger // warning handler
	)

	Warning = log.New(
		os.Stdout,
		color.YellowString("WARNING")+":",
		log.Ldate|log.Ltime,
	)

	Warning.Println(err)
}

// WarningReceiver reports warning on stdout via a blocking channel
func WarningReceiver(err <-chan error) {
	for {
		Fatal(<-err)
	}
}
