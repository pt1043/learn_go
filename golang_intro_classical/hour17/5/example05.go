package main

import (
	"flag"
	"fmt"
	"os"
)

func flagUsage() {
	usageText := `example05 is an example cli tool.

Usage:
example05 command [arguments]
The commands are:
uppercase uppercase a string
lowercase lowercase a string
Use "example05 [command] --help" for more information about a command.`
	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func main() {

	flag.Usage = flagUsage
	_ = flag.NewFlagSet("uppercase", flag.ExitOnError)
	_ = flag.NewFlagSet("lowercase", flag.ExitOnError)

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}
}
