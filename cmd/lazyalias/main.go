package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sergiorivas/lazyalias/internal/core"
)

var (
	version     = "v0.1.11"
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
	err := commander.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
