package GoError

import (
	// standard packages
	"log" // simple logging package
	"os"  // platform-independent interface to operating system functionality

	// external packages
	"github.com/fatih/color" // colorized outputs in terms of ANSI Escape Codes
)

// Warning ...
// -------
// reports warning on stdout
//
// input: err - error output
//
// return:
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
