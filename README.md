# brewer

CLI app to update Homebrew formulae.

## Installation

To get the most up to date binaries, check [the releases][releases] for the
pre-built binary for your system.

[releases]: https://github.com/tombell/brewer/releases

You can also `go get` to install from source.

    go get -u github.com/tombell/brewer/cmd/brewer

## Usage

Use the `-h/--help` flag to see all the available flags when running **brewer**.

You will need a [GitHub API token][api-token] when running **brewer**. It's
advised you create a token specifically for **brewer**.

[api-token]: https://github.com/settings/tokens

You can update four different parts of a formula with **brewer**.

- the tag as part of the URL `:tag => "..."` with `--tag`
- the revision as part of the URL `:revision => "..."` with `--revision`
- the URL `url "..."` with `--url`
- the SHA256 `sha256 "..."` with `--sha256`

The path in the repository to the formula to update is passed as the `--formula`
flag.

You also have to specify the repository owner and name (the name defaults to
`homebrew-formulae` if omitted).

The commit message, author, and author's email must be given for the commit that
updates the formula.

    brewer --token $GITHUB_TOKEN \
           --owner tombell \
           --name homebrew-formulae \
           --formula Formula/brewer.rb \
           --tag v1.0.0 \
           --revision beea04da6f030de641408ca265e31fb01e7dc22f \
           --commit-message "Update Formula/brewer.rb to v1.0.0" \
           --commit-author "Tom Bell" \
           --commit-email "tomb@tomb.io"

This will pull down `Formula/brewer.rb` from the repository, and update the tag
and revision lines with `v1.0.0` and `beea04da6f030de641408ca265e31fb01e7dc22f`.
With the commit message `Update Formula/brewer.rb to v1.0.0`, by Tom Bell.
