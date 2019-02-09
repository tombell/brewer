package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tombell/brewer"
)

const helpText = `usage: brewer [options]

Special options:
  --help     show this message, then exit
  --version  show the brewer version number, then exit

GitHub options:
  --token  GitHub API token
  --owner  GitHub repository owner of the Homebrew tap
  --repo   GitHub repository name of the Homebrew tap (defaults to homebrew-formulae)

Formula options:
  --formula   Path to the formula in the git repository
  --tag       New git tag to update in the formula
  --revision  New revision SHA to update in the formula
  --url       New URL to update in the formula
  --sha256    New SHA256 to update in the formula

Commit options:
  --commit-message  Commit message for the formula update
  --commit-author   Commit author for the formula update
  --commit-email    Commit email for the formula update

`

var (
	vrsn = flag.Bool("version", false, "")

	token = flag.String("token", "", "")
	owner = flag.String("owner", "", "")
	name  = flag.String("name", "homebrew-formulae", "")

	formula  = flag.String("formula", "", "")
	tag      = flag.String("tag", "", "")
	revision = flag.String("revision", "", "")
	url      = flag.String("url", "", "")
	sha256   = flag.String("sha256", "", "")

	message = flag.String("commit-message", "", "")
	author  = flag.String("commit-author", "", "")
	email   = flag.String("commit-email", "", "")
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
		exit("must provide --token flag for GitHub API token")
	}

	if *owner == "" {
		exit("must provide --owner flag for GitHub repository owner")
	}

	if *formula == "" {
		exit("must provide --formula flag for formula path")
	}

	if *author == "" {
		exit("must provide --author flag for update commit author")
	}

	if *email == "" {
		exit("must provide --email flag for update commit email")
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *vrsn {
		fmt.Fprintf(os.Stdout, "brewer %s (%s)\n", version, commit)
		os.Exit(0)
	}

	validateFlags()

	config := brewer.Config{
		Token:         *token,
		Owner:         *owner,
		Repo:          *name,
		Path:          *formula,
		Tag:           *tag,
		Revision:      *revision,
		URL:           *url,
		SHA256:        *sha256,
		CommitMessage: *message,
		CommitAuthor:  *author,
		CommitEmail:   *email,
	}

	if err := brewer.Run(config); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
