package brewer

import (
	"context"
	"fmt"

	"github.com/google/go-github/v18/github"
	"golang.org/x/oauth2"
)

// Config represents the command line flags passed into the Run function.
type Config struct {
	Token string
	Owner string
	Repo  string

	Path     string
	Tag      string
	Revision string
	URL      string
	SHA256   string

	CommitMessage string
	CommitAuthor  string
	CommitEmail   string
}

// Run performs the updating of the given Homebrew formula.
func Run(config Config) error {
	fmt.Printf("Updating formula (%s)...\n", config.Path)

	client := createGitHubClient(config.Token)

	formula, err := fetchFormula(client, config)
	if err != nil {
		return err
	}

	if config.Tag != "" {
		fmt.Printf("- Updating tag to %s...\n", config.Tag)
		formula.UpdateTag(config.Tag)
	}

	if config.Revision != "" {
		fmt.Printf("- Updating revision to %s...\n", config.Revision)
		formula.UpdateRevision(config.Revision)
	}

	if config.URL != "" {
		fmt.Printf("- Updating URL to %s...\n", config.URL)
		formula.UpdateURL(config.URL)
	}

	if config.SHA256 != "" {
		fmt.Printf("- Updating SHA256 to %s...\n", config.SHA256)
		formula.UpdateSHA256(config.SHA256)
	}

	if err := updateFormula(client, config, formula); err != nil {
		return err
	}

	fmt.Printf("Finished updating formula (%s)\n", config.Path)

	return nil
}

func createGitHubClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)
	return github.NewClient(tc)
}

func fetchFormula(c *github.Client, config Config) (*Formula, error) {
	opt := &github.RepositoryContentGetOptions{}

	file, _, _, err := c.Repositories.GetContents(context.Background(), config.Owner, config.Repo, config.Path, opt)
	if err != nil {
		return nil, err
	}

	contents, err := file.GetContent()
	if err != nil {
		return nil, err
	}

	return &Formula{Path: *file.Path, Contents: contents, FileSHA: *file.SHA}, nil
}

func updateFormula(c *github.Client, config Config, formula *Formula) error {
	opt := &github.RepositoryContentFileOptions{
		Content: []byte(formula.Contents),
		SHA:     github.String(formula.FileSHA),

		Message: github.String(config.CommitMessage),
		Committer: &github.CommitAuthor{
			Name:  github.String(config.CommitAuthor),
			Email: github.String(config.CommitEmail),
		},
	}

	_, _, err := c.Repositories.UpdateFile(context.Background(), config.Owner, config.Repo, formula.Path, opt)
	if err != nil {
		return err
	}

	return nil
}
