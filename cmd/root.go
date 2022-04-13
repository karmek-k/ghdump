package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v43/github"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ghdump <username>",
	Short: "Bulk clone GitHub repos for backup",
	Long: `ghdump allows you to automatically clone your (or someone else's)
repositories.

Issuing ghdump without any subcommand will show some info
about a user.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client := github.NewClient(nil)
		username := args[0]

		fmt.Printf("User %s\n", username)
		
		const MAX_REPOS = 100

		opt := &github.RepositoryListOptions{
			ListOptions: github.ListOptions{PerPage: MAX_REPOS},
		}
		repos, _, err := client.Repositories.List(ctx, username, opt)
		if err != nil {
			return err
		}
		
		count := len(repos)
		fmt.Printf("Repo count (public): %d", count)
		if count >= MAX_REPOS {
			fmt.Print(" or more")
		}

		fmt.Println()

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ghdump.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
