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
	return "", nil
}

// UpdateFormula updates the contents of the formula from the given repository.
func UpdateFormula(c *github.Client, owner, repo, formula, contents string) error {
	// TODO: implement
	return nil
}
