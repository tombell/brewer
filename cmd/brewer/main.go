package main

import (
	"flag"
	"fmt"
	"os"
)

const helpText = `usage: brewer [options]

`

var (
	version = flag.Bool("version", false, "")

	token = flag.String("token", "", "")
	owner = flag.String("owner", "", "")
	name  = flag.String("name", "homebrew-formulae", "")

	formula = flag.String("formula", "", "")
)

func usage() {
	fmt.Fprintf(os.Stderr, helpText)
	os.Exit(2)
}

func validateFlags() {
	if *token == "" {
		fmt.Fprintf(os.Stderr, "must provide -token flag\n")
		os.Exit(1)
	}

	if *owner == "" || *name == "" {
		fmt.Fprintf(os.Stderr, "must provide repo owner and name\n")
		os.Exit(1)
	}

	if *formula == "" {
		fmt.Fprintf(os.Stderr, "must provide formula name\n")
		os.Exit(1)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stdout, "brewer %s (%s)\n", Version, Commit)
		os.Exit(0)
	}

	validateFlags()
}
