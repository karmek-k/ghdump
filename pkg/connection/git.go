package connection

import (
	"github.com/go-git/go-git/v5"
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/go-github/v43/github"
)

type CloneRepoOptions struct {
	Repo      *github.Repository
	CloneBare bool
	Directory string
	GitAuth   *gitHttp.BasicAuth
}

func GitBasicAuth(username, password string) *gitHttp.BasicAuth {
	if password != "" {
		return &gitHttp.BasicAuth{
			Username: username,
			Password: password,
		}
	}

	return nil
}

func CloneRepo(options CloneRepoOptions) {
	git.PlainClone(
		options.Directory+"/"+*options.Repo.Name,
		options.CloneBare,
		&git.CloneOptions{
			URL:  *options.Repo.CloneURL,
			Auth: options.GitAuth,
		},
	)
}
