/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/karmek-k/ghdump/pkg/connection"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone <username>",
	Short: "Clones user's (or organization's) repositories",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		username := args[0]
		pat := cmd.Flag("token").Value.String()
		visibility := cmd.Flag("visibility").Value.String()

		client := connection.CreateGitHubClient(ctx, username, pat)
		gitAuth := connection.GitBasicAuth(username, pat)

		repos, err := client.ListUserRepos(ctx, visibility)
		if err != nil {
			return err
		}

		cloneForks, err := cmd.Flags().GetBool("clone-forks")
		if err != nil {
			return err
		}

		for _, repo := range repos {
			// don't include orgs' repos
			if repo.Owner.GetLogin() != username {
				continue
			}

			// don't clone forks if the user doesn't want to
			if *repo.Fork && !cloneForks {
				continue
			}

			fmt.Printf("Cloning %s...\n", *repo.Name)

			connection.CloneRepo(connection.CloneRepoOptions{
				Repo:      repo,
				CloneBare: false,
				Directory: cmd.Flag("output-dir").Value.String(),
				GitAuth:   gitAuth,
			})
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
	cloneCmd.Flags().BoolP("clone-forks", "f", false, "Whether forks should be cloned")
	cloneCmd.Flags().StringP("token", "t", "", "Personal access token (PAT)")
	cloneCmd.Flags().StringP("output-dir", "o", "dump", "The directory repos should be cloned to")
	cloneCmd.Flags().StringP("visibility", "v", "all", "Repo visibility (all, public, private)")
}
