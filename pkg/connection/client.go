package connection

import (
	"context"
	"net/http"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	handle   *github.Client
	Username string
	Pat      string
}

func CreateGitHubClient(ctx context.Context, username, pat string) *GitHubClient {
	var httpClient *http.Client = nil
	if pat != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: pat},
		)
		httpClient = oauth2.NewClient(ctx, ts)
	}

	return &GitHubClient{
		handle: github.NewClient(httpClient),
		Pat:    pat,
	}
}

func ListOptions(visibility string) *github.RepositoryListOptions {
	return &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
		Visibility:  visibility,
	}
}

func (c *GitHubClient) ListUserRepos(ctx context.Context, visibility string) ([]*github.Repository, error) {
	var result []*github.Repository
	opt := ListOptions(visibility)

	for {
		repos, resp, err := c.handle.Repositories.List(ctx, c.Username, opt)
		if err != nil {
			return nil, err
		}

		result = append(result, repos...)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return result, nil
}
