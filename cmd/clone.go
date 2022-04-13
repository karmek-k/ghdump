/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v43/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone <username>",
	Short: "Clones user's repositories",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: too much code :)
		ctx := context.Background()

		username := args[0]

		var httpClient *http.Client = nil
		pat := cmd.Flag("token").Value.String()
		if pat != "" {
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: pat},
			)
			httpClient = oauth2.NewClient(ctx, ts)
		}

		client := github.NewClient(httpClient)

		opt := &github.RepositoryListOptions{
			ListOptions: github.ListOptions{PerPage: 10},
			Visibility: cmd.Flag("visibility").Value.String(),
		}

		for {
			repos, resp, err := client.Repositories.List(ctx, username, opt)
			if err != nil {
				return err
			}

			for _, repo := range repos {
				// process repos

				// don't include orgs' repos
				ownerLogin := repo.Owner.GetLogin()
				if ownerLogin != username {
					continue
				}

				fmt.Printf("%s (%s)\n", *repo.Name, *repo.GitURL)
			}
			
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	cloneCmd.Flags().StringP("token", "t", "", "Personal access token (PAT)")
	cloneCmd.Flags().StringP("visibility", "v", "all", "Repo visibility (all, public, private)")
}
