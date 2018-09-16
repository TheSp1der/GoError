package goerror

import (
	// standard packages
	"log" // simple logging package
	"os"  // platform-independent interface to operating system functionality

	// external packages
	"github.com/fatih/color" // colorized outputs in terms of ANSI Escape Codes
)

// Fatal ...
// -------
// reports error on stdout and halts the program
//
// input: err - error output
//
// return:
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
