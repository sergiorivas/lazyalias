package main

import (
	"flag"
	"fmt"

	"github.com/sergiorivas/lazyalias/internal/core"
)

var (
	version     = "v0.1.5"
	showVersion = flag.Bool("version", false, "show version information")
)

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Printf("🤓 lazyalias version %s\n", version)
		return
	}

	fmt.Printf("Welcome to LAZYALIAS 🎉🎉🎉\n")
	commander := core.NewCommander()
	commander.Run()
}
