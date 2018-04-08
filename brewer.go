package brewer

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// CreateGitHubClient creates a new GitHub API client with a given API token for
// authentication.
func CreateGitHubClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)
	return github.NewClient(tc)
}

// FetchFormula fetches the contents of the formula from the given repository.
func FetchFormula(c *github.Client, owner, repo, formula string) (string, error) {
	// TODO: implement
	// get formula file contents...
	// - https://github.com/google/go-github/blob/master/github/repos_contents.go#L143
	return "", nil
}

// UpdateFormula updates the contents of the formula from the given repository.
func UpdateFormula(c *github.Client, owner, repo, formula, contents string) error {
	// TODO: implement
	// update formula file contents...
	// - https://github.com/google/go-github/blob/master/github/repos_contents.go#L192
	return nil
}
