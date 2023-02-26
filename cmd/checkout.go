/*
Copyright © 2023 Darren Yim <darrenyxj@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"gitx/pkg/git"
)

// checkoutCmd represents the Checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout [arg:branch name]",
	Args:  cobra.ExactArgs(1),
	Short: "Checkout specified branch from first level of directories",
	Long: `Performs the following actions to all directories in the current folder
1. stashes current work
2. checks out branch specified`,
	Run: runCheckout,
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func runCheckout(cmd *cobra.Command, args []string) {
	processor := git.FolderProcessor{
		ProcessFn: func(cmd *cobra.Command, args []string) {
			git.Stash()
			git.Checkout(args[0])
		},
	}

	processor.ProcessFolder(cmd, args)
}
