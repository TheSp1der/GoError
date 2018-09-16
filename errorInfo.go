package goerror

import (
	// standard packages
	"log" // simple logging package
	"os"  // platform-independent interface to operating system functionality

	// external packages
	"github.com/fatih/color" // colorized outputs in terms of ANSI Escape Codes
)

// Info ...
// -------
// reports info on stdout
//
// input: err - error output
//
// return:
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
