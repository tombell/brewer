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

func exit(msg string) {
	fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(1)
}

func validateFlags() {
	if *token == "" {
		exit("must provide -token flag")
	}

	if *owner == "" || *name == "" {
		exit("must provide repo owner and name")
	}

	if *formula == "" {
		exit("must provide formula name")
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
